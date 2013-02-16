package boop

import (
	"fmt"
	"net/http"
	"os"
	"os/exec"
	"strconv"
	"strings"
        "time"
)

var gConfig *Config

const NormalReturn = 0
const ConfigError = 1
const ListenError = 2

func BoopMain() {
	// Try to load the gConfig file
	var err error
	gConfig, err = LoadConfig()
	// If it doesn't load shut down
	if err != nil {
		fmt.Println("Error loading configuration file: " + err.Error())
		fmt.Println("Shutting down due gConfig error")
		os.Exit(ConfigError)
	}
	fmt.Printf("Got gConfig %#v\n", *gConfig)
	fmt.Printf("Starting http server on port %d\n", gConfig.Port)
	// Register our handler to handle all communications
	http.Handle("/", http.HandlerFunc(httpRequestHandler))
	fmt.Println("Listening...")
	err = http.ListenAndServe("0.0.0.0:"+strconv.Itoa(gConfig.Port), nil)
	if err != nil {
		fmt.Println("ListenAndServ Error: ", err)
		os.Exit(ListenError)
	}
}

func httpRequestHandler(w http.ResponseWriter, req *http.Request) {
	actionPath := req.Method + " " + req.URL.Path
	fmt.Println("Received request for " + actionPath + " from " + requestIpAddress(req))
	command := commandForActionPath(actionPath)
	if command == nil {
		w.WriteHeader(404)
		fmt.Println("No associated command.  Returning 404 (Not Found).")
		return
	}

        command.Lock()
        defer command.Unlock()

	if !authorized(req, command) {
		w.WriteHeader(401)
		fmt.Println("Command is valid, but user is unauthorized.  Returning 401 (Unauthorized)")
		return
	}

        if !command.MeetsRateLimit() {
                w.WriteHeader(429)
                fmt.Println("Command is valid, but would violate the rate limit.  Returning 429 (Too Many Requests)")
                return
        }

        command.Execute()
        w.WriteHeader(200)
        fmt.Printf("Command %v executed\n", command);
}

func commandForActionPath(actionPath string) *Command {
	// Just in case no commands have been defined

	for k, v := range gConfig.Commands {
		if v.Path == actionPath {
			return &gConfig.Commands[k]
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
	ipAddr := requestIpAddress(req)

	// List of lists that we should check.  The order determines precedence, though
	// if an earlier list is empty it is ignored.
	authLists := [...][]string{command.OnlyAllowIps, gConfig.OnlyAllowIps}

	for _, authList := range authLists {
		if len(authList) != 0 {
			for _, v := range authList {
				if v == ipAddr {
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

func (c *Command) Execute() {
	cmd := exec.Command(os.Getenv("SHELL"), "-c", c.Command)
        (*c).LastTimeRun = time.Now()
	go cmd.Run()
}

func (c *Command) MeetsRateLimit() bool {
        if c.LimitRate == 0 {
          return true
        }

        if time.Now().After(c.LastTimeRun.Add(time.Duration(c.LimitRate) * time.Second)) {
          return true
        }

        return false
}
