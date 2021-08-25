package main

import (
	"encoding/json"
	"log"
	"os"
	"os/signal"
	"runtime"

	"app/route"
	"app/shared/database"
	"app/shared/email"
	"app/shared/jsonconfig"
	"app/shared/recaptcha"
	"app/shared/server"
	"app/shared/session"
	"app/shared/view"
	"app/shared/view/plugin"

	"unsafe"

	"github.com/zserge/lorca"
)

// #include <stdio.h>
// #include <stdlib.h>
//
// static void myprint(char* s) {
//   printf("%s\n", s);
// }
import "C"

// *****************************************************************************
// Application Logic
// *****************************************************************************
var ui lorca.UI
var err error

func init() {
	// Verbose logging with file name and line number
	log.SetFlags(log.Lshortfile)

	// Use all CPU cores
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func main() {
	cs := C.CString("Hello from stdio")
	C.myprint(cs)
	C.free(unsafe.Pointer(cs))
	// Load the configuration file
	jsonconfig.Load("config"+string(os.PathSeparator)+"config.json", config)

	// Configure the session cookie store
	session.Configure(config.Session)

	// Connect to database
	database.Connect(config.Database)

	// Configure the Google reCAPTCHA prior to loading view plugins
	recaptcha.Configure(config.Recaptcha)

	// Setup the views
	view.Configure(config.View)
	view.LoadTemplates(config.Template.Root, config.Template.Children)
	view.LoadPlugins(
		plugin.TagHelper(config.View),
		plugin.NoEscape(),
		plugin.PrettyTime(),
		recaptcha.Plugin())

	go loadUI()
	// Start the listener
	server.Run(route.LoadHTTP(), route.LoadHTTPS(), config.Server)

}

// *****************************************************************************
// Application Settings
// *****************************************************************************

// config the settings variable
var config = &configuration{}

// configuration contains the application settings
type configuration struct {
	Database  database.Info   `json:"Database"`
	Email     email.SMTPInfo  `json:"Email"`
	Recaptcha recaptcha.Info  `json:"Recaptcha"`
	Server    server.Server   `json:"Server"`
	Session   session.Session `json:"Session"`
	Template  view.Template   `json:"Template"`
	View      view.View       `json:"View"`
}

// ParseJSON unmarshals bytes to structs
func (c *configuration) ParseJSON(b []byte) error {
	return json.Unmarshal(b, &c)
}

func loadUI() {
	//Possible Sleep time.Sleep(time.Second * 0)

	args := []string{}
	if runtime.GOOS == "linux" {
		args = append(args, "--class=Lorca")
	}
	ui, err = lorca.New("", "", 1280, 900, args...)
	if err != nil {
		log.Fatal(err)
	}

	ui.Load("http://localhost:8080/")
	defer ui.Close()

	// Wait until the interrupt signal arrives or browser window is closed
	sigc := make(chan os.Signal)
	signal.Notify(sigc, os.Interrupt)
	select {
	case <-sigc:
	case <-ui.Done():
	}
	log.Println("exiting...")
	os.Exit(0)
}
