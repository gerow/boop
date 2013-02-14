package boop

import (
	"fmt"
	"net/http"
	"os"
	//"os/exec"
	"strconv"
	"strings"
)

var config *Config

const NORMAL_RETURN = 0
const CONFIG_ERROR = 1
const LISTEN_ERROR = 2

func BoopMain() {
	// Try to load the config file
	var err error
	config, err = LoadConfig()
	// If it doesn't load shut down
	if err != nil {
		fmt.Println("Error loading configuration file: " + err.Error())
		fmt.Println("Shutting down due config error")
		os.Exit(CONFIG_ERROR)
	}
	fmt.Printf("Got config %#v\n", *config)
	fmt.Printf("Starting http server on port %d\n", config.Port)
	// Register our handler to handle all communications
	http.Handle("/", http.HandlerFunc(httpRequestHandler))
	err = http.ListenAndServe("0.0.0.0:"+strconv.Itoa(config.Port), nil)
	fmt.Println("Listening...")
	if err != nil {
		fmt.Println("ListenAndServ Error: ", err)
		os.Exit(LISTEN_ERROR)
	}
}

func httpRequestHandler(w http.ResponseWriter, req *http.Request) {
	action_path := req.Method + " " + req.URL.Path
	fmt.Println("Received request for " + action_path + " from " + requestIpAddress(req))
	command := commandForActionPath(action_path)
	if command == nil {
		w.WriteHeader(404)
		fmt.Println("No associated command.  Returning 404 (Not Found).")
		return
	}

	if !authorized(req, command) {
		w.WriteHeader(401)
		fmt.Println("Command is valid, but user is unauthorized.  Returning 401 (Unauthorized)")
		return
	}
	//cmd := exec.Command(os.Getenv("SHELL"), "-c", v)
	//go cmd.Run()
	//w.WriteHeader(200)
}

func commandForActionPath(action_path string) *Command {
	// Just in case no commands have been defined

	for _, v := range config.Commands {
		if v.Path == action_path {
			return &v
		}
	}

	return nil
}

func requestIpAddress(req *http.Request) string {
	// Should probably do some error checking here...
	return strings.Split(req.RemoteAddr, ":")[0]
}

// We verify authenticity by first looking on the command's restricted
// ips.  We then look at the global one of there is none for the command.
func authorized(req *http.Request, command *Command) bool {
	ip_addr := requestIpAddress(req)

	// List of lists that we should check.  The order determines precedence, though
	// if an earlier list is empty it is ignored.
	authentication_lists := [...][]string{command.OnlyAllowIps, config.OnlyAllowIps}

	for _, auth_list := range authentication_lists {
		if len(auth_list) != 0 {
			for _, v := range auth_list {
				if v == ip_addr {
					// Found it in the lest, authorized
					return true
				}
			}
			// Couldn't find it in the list, unauthorized
			return false
		}
	}

	// All of the lists were empty, so this server is open to
	// connections from anyone
	return true
}
