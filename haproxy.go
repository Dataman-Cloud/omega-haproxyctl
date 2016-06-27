package main

import (
	"encoding/json"
	"flag"
	"io"
	"log"
	"net/http"
	"os"
	"os/exec"
	"os/signal"
	"syscall"
	"time"

	"github.com/Dataman-Cloud/HAServer/Godeps/_workspace/src/github.com/go-martini/martini"
	"github.com/Dataman-Cloud/HAServer/Godeps/_workspace/src/github.com/natefinch/lumberjack"
	"github.com/Dataman-Cloud/HAServer/cmd"
	"github.com/Dataman-Cloud/HAServer/configuration"
)

/*
	Commandline arguments
*/
var (
	logPath        string
	serverBindPort string
	runtime        cmd.Runtime
	reloadChan     chan int
)

func init() {
	flag.StringVar(&logPath, "log", "", "Log path to a file. Default logs to stdout")
	flag.StringVar(&serverBindPort, "bind", ":5004", "Bind HTTP server to a specific port")
	reloadChan = make(chan int, 3)
}

type Response struct {
	Code int    `json:"code"`
	Err  string `json:"err"`
}

func main() {
	flag.Parse()
	configureLog()

	// Load configuration
	conf := configuration.Configs()
	runtime = cmd.Runtime{
		Binary:   conf.HAProxy.Command,
		SockFile: conf.HAProxy.SockFile,
	}

	// Wait for died children to avoid zombies
	signalChannel := make(chan os.Signal, 2)
	signal.Notify(signalChannel, os.Interrupt, syscall.SIGCHLD)
	go func() {
		for {
			sig := <-signalChannel
			if sig == syscall.SIGCHLD {
				r := syscall.Rusage{}
				syscall.Wait4(-1, nil, 0, &r)
			}
		}
	}()

	// Handle gracefully exit
	registerOSSignals()

	// Start server
	initServer(&conf)
}

func initServer(conf *configuration.Configuration) {
	// begin listen to reload haproxy
	go ReloadHaproxyConfig(conf)
	// Status live information
	router := martini.Classic()
	// API
	router.Group("/api", func(api martini.Router) {
		// State API
		api.Get("/status", HealthCheck)
		// Service API
		api.Put("/haproxy", servicesApi)
		// Weight API
		api.Put("/weight", updateWeight)
	})

	router.RunOnAddr(serverBindPort)
}

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	conf := configuration.Configs()
	if err := validateConfig(&conf); err != nil {
		log.Println("validateConfig got error: ", err)
		http.Error(w, "Failed to validate haproxy config", http.StatusInternalServerError)
		return
	}

	io.WriteString(w, "Successed to validate haproxy.cfg")
	return
}

func updateWeight(w http.ResponseWriter, r *http.Request) {
	servers := []struct {
		Backend string `param:"backend" json:"backend"`
		Server  string `param:"server" json:"server"`
		Weight  int    `param:"weight" json:"weight"`
	}{}
	decoder := json.NewDecoder(r.Body)
	defer r.Body.Close()
	err := decoder.Decode(&servers)
	if err != nil {
		log.Println("Error: cannot parse server weight", err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if len(servers) == 0 {
		log.Println("empty servers")
		responseJSON(w, Response{Code: 0})
		return
	}

	for _, server := range servers {
		log.Println("setting weight", server.Backend, server.Server, server.Weight)
		out, err := runtime.SetWeight(server.Backend, server.Server, server.Weight)
		if err != nil {
			log.Println("Error: cannot set server weight", err.Error())
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		log.Println("set weight", out)
	}

	responseJSON(w, Response{Code: 0})
}

func servicesApi(w http.ResponseWriter, r *http.Request) {
	if len(reloadChan) < 3 {
		reloadChan <- 1
		responseJSON(w, Response{Code: 0, Err: ""})
		return
	}

	log.Println("HAProxy reload queue is fully, ignore this message")
	responseJSON(w, Response{Code: 0, Err: "reloading, ignore message"})
}

func ReloadHaproxyConfig(conf *configuration.Configuration) {
	for {
		select {
		case <-reloadChan:
			reloadConfig(conf)
			time.Sleep(time.Second * 1)
		}
	}
}

// validate haproxy config
func validateConfig(conf *configuration.Configuration) error {
	log.Println("Validating config")
	err := execCommand(conf.HAProxy.ReloadValidationCommand)
	if err != nil {
		log.Println("Validat config Error: ", err.Error())
		return err
	}

	return nil
}

// reload and update haproxy config
func reloadConfig(conf *configuration.Configuration) (bool, error) {
	// delete drop tcp sync message
	defer func() {
		log.Println("After reload")
		if err := execCommand(conf.HAProxy.AfterReload); err != nil {
			log.Println("WARN: AfterReload Command got error: ", err.Error())
		}
	}()

	err := validateConfig(conf)
	if err != nil {
		return false, err
	}

	log.Println("Before reload")
	// add drop tcp sync message
	if err := execCommand(conf.HAProxy.BeforeReload); err != nil {
		log.Println("WARN: BeforeReload Command got error: ", err.Error())
	}

	log.Println("Reload config")
	if err := execCommand(conf.HAProxy.ReloadCommand); err != nil {
		log.Println("Reload config Error: ", err.Error())
		return false, err
	}

	return true, nil
}

func execCommand(cmd string) error {
	log.Printf("Exec cmd: %s \n", cmd)
	output, err := exec.Command("sh", "-c", cmd).CombinedOutput()
	if err != nil {
		log.Println(err.Error())
		log.Println("Output:\n" + string(output[:]))
	}
	return err
}

func responseJSON(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	bites, _ := json.Marshal(data)
	w.Write(bites)
}

func configureLog() {
	if len(logPath) > 0 {
		log.SetOutput(io.MultiWriter(&lumberjack.Logger{
			Filename: logPath,
			// megabytes
			MaxSize:    100,
			MaxBackups: 3,
			//days
			MaxAge: 28,
		}, os.Stdout))
	}
}

func registerOSSignals() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		for _ = range c {
			log.Println("Server Stopped")
			os.Exit(0)
		}
	}()
}
