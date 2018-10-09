package main

import (
    "html/template"
    "net/http"
    "net"
    "log"
)

type PageVariables struct {
	Ip         string
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
    http.HandleFunc("/", IndexPage)
    log.Fatal(http.ListenAndServe(":8080", nil))
}
