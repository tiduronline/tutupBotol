package main 

import (
    "strings"
    "io/ioutil"
    "fmt"
    "runtime"
    "os"
)

// Simple Config using struct
// ===================================
type Config struct {
    BofHeader string
    EofFooter string
    UpdateEndpoint string
    HostFile string
    Debug bool
}

func init_config() Config {
    config := Config{}
    config.BofHeader        = "# BOF D COINS"
    config.EofFooter        = "# EOF D COINS"
    config.UpdateEndpoint   = "https://raw.githubusercontent.com/hoshsadiq/adblock-nocoin-list/master/hosts.txt"
    config.HostFile         = "/etc/hosts"

    if runtime.GOOS=="windows" {
        config.HostFile = "C:\\Windows\\System32\\drivers\\etc\\hosts"
    }
    
    config.Debug            = false

    if _, err := os.Stat(config.HostFile); os.IsNotExist(err) {
        fmt.Println(err)
        fmt.Println("Hosts file is not found")
        os.Exit(0)
    }

    return config
}
// End Of Config
// ===================================


// Easy Error Checker
// with this we can set if error is verbose 
// or just quite log with Debug mode
func err(e error) {
    config := init_config()
    if e != nil {
        if config.Debug {
            panic(e)
        } else {
            fmt.Println(e)
        }
    }
}


func main() {
    config  := init_config()
    option  := parseArgs()

    var hosts []string
    content, e := ioutil.ReadFile(config.HostFile)
    err(e)

    s_content := string(content)
    l_content := strings.Split(s_content, "\n")

    for _, host := range l_content {
        hosts = append(hosts, string(host))
    }

    if option == "block" || option == "update" {
        updateList(config, hosts)
        fmt.Println("Hosts has been updated")
    } else if option == "clean" {
        cleanList(config, hosts)
        fmt.Println("Successfull cleaned")
    }
}
