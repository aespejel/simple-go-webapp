package main

import (
    "html/template"
    "net/http"
    "net"
    "log"
    "time"
    "os"
    "fmt"
    "strconv"
)

type PageVariables struct {
	Ip         string
}

func DelayStartup() {
    if startup_delay, ok := os.LookupEnv("STARTUP_DELAY"); ok {
        startup_delay_int, err := strconv.Atoi(startup_delay)
        if err != nil {
            panic(fmt.Sprintf("STARTUP_DELAY environmental variable is not a number"))
        }
        time.Sleep(time.Duration(startup_delay_int) * time.Second)
    }
}

func IndexPage(w http.ResponseWriter, r *http.Request){
    IndexPageVars := PageVariables{ //store the page variables in a struct
      Ip: getMyIp(),
    }
    t, err := template.ParseFiles("templates/index.html")
    err = t.Execute(w, IndexPageVars)
    if err != nil {
	  log.Print("template executing error: ", err)
	}
}

func getMyIp() string {
    addrs, err := net.InterfaceAddrs()
    if err != nil {
        panic(err)
    }
    for _, address := range addrs {
        // if the address is not loopback, return it
        if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
            if ipnet.IP.To4() != nil {
                return ipnet.IP.String()
            }
        }
    }
    return ""
}

func main() {
    DelayStartup()
    http.HandleFunc("/", IndexPage)
    log.Fatal(http.ListenAndServe(":8080", nil))
}
