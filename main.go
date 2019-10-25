package main

import (
	"exp_with_go_lang/lib"
	"flag"
	"fmt"
	util "github.com/eagle7410/go_util/libs"
	"log"
	"net"
	"net/http"
	"os"
	"os/exec"
	"syscall"
	"time"
)

const port = ":8080"

var (
	isInstallService, isUninstallService bool
)

func init() {

	util.OpenLogFile()

	if err := lib.ENV.Init(); err != nil {
		util.LogFatalf("Error on initializing environment : %s", err)
	}

	util.Env.SetEnv(&lib.ENV)
	flag.BoolVar(&isInstallService, "install", false, "Установить как linux service")
	flag.BoolVar(&isUninstallService, "uninstall", false, "Деинсталяция linux service")
}

func isServiceAction () bool {
	if isInstallService {
		lib.InstallAsService()
		return true
	}

	if isUninstallService {
		lib.UninstallService()
		return true
	}

	return false
}

func main() {

	flag.Parse()

	if isServiceAction() {
		util.Logf("\n !!!App use service action \n")
		return
	}

	router := lib.GetRouter()
	router.HandleFunc("/grace", GraceHandler)

	middleware := util.SetCorsMiddleware(
		util.LogRequest(
			router))

	srv := &http.Server{
		Addr:         "0.0.0.0" + port,
		// Good practice to set timeouts to avoid S
		// Good practice to set timeouts to avoid Slowloris attacks.
		WriteTimeout: time.Minute * 10,
		ReadTimeout:  time.Minute * 10,
		IdleTimeout:  time.Minute * 10,
		Handler: middleware, // Pass our instance of gorilla/mux in.
	}

	// Run our server in a goroutine so that it doesn't block.

	util.LogAppRun(port)
	var err error

	if *FD != 0 {
		util.Logf("Starting with FD %v", *FD)
		file1 = os.NewFile(uintptr(*FD), "parent socket")
		listener1, err = net.FileListener(file1)
		if err != nil {
			util.LogFatalf("fd listener failed: %v", err)
		}
	} else {
		util.Logf("Virgin Start")

		listener1, err = net.Listen("tcp", srv.Addr)

		if err != nil {
			util.LogFatalf("listener failed: %v \n", err)
		}
	}
	err = srv.Serve(listener1)
	util.LogEF("err 97 %v", err)
	util.Logf("EXITING", PID)
	<-exit1
	util.Logf("EXIT %v", PID)

	//c := make(chan os.Signal, 1)
	//// We'll accept graceful shutdowns when quit via SIGINT (Ctrl+C)
	//// SIGKILL, SIGQUIT or SIGTERM (Ctrl+/) will not be caught.
	//signal.Notify(c, os.Interrupt)
	//
	//// Block until we receive our signal.
	//<-c
	//
	//// Create a deadline to wait for.
	//ctx, cancel := context.WithTimeout(context.Background(), time.Second * 30 )
	//defer cancel()
	//// Doesn't block if no connections, but will otherwise wait
	//// until the timeout deadline.
	//srv.Shutdown(ctx)
	//// Optionally, you could run srv.Shutdown in a goroutine and block on
	//// <-ctx.Done() if your application should wait for other services
	//// to finalize based on context cancellation.
	//util.Logf("shutting down")
	//os.Exit(0)
}
var FD *int = flag.Int("fd", 0, "Server socket FD")
var PID int = syscall.Getpid()
var listener1 net.Listener
var file1 *os.File = nil
var exit1 chan int = make(chan int)
var stop1 = false
func GraceHandler(w http.ResponseWriter, req *http.Request) {
	util.Logf("GraceHandler %v", req.Method)
	if(stop1){
		fmt.Fprintf(w, "stopped %d %s", PID, time.Now().String())
	}
	stop1 = true
	fmt.Fprintf(w, "grace %d %s", PID, time.Now().String())

	go func() {
		defer func() { util.Logf("GoodBye") }()
		listener2 := listener1.(*net.TCPListener)
		file2, err := listener2.File()
		if err != nil {
			util.LogEF("err is 141 %v", err)
		}
		fd1 := int(file2.Fd())

		fd2, err := syscall.Dup(fd1)
		if err != nil {
			log.Fatalln("Dup error:", err)
		}

		listener1.Close()
		if file1 != nil {
			file1.Close()
		}

		cmd := exec.Command("./serve", fmt.Sprint("-fd=", fd2))
		err = cmd.Start()
		if err != nil {
			log.Fatalln("grace starting error:", err)
		}

		log.Println("sleep11", PID)
		time.Sleep(10 * time.Second)
		log.Println("exit after sleep", PID)
		exit1<-1
	}()
}
