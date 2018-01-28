package main

import (
    "strings"
    "fmt"
    "net/http"
    "io/ioutil"
    "flag"
    "regexp"
)


// Argument Parser for golang command line
// you only have one option to passed as argument
//
// Availabel option:
//     -block = Block domain of coins
//     -clean = Unblock coin domain
//     -update = Update list of coin domain
func parseArgs() string {

    choiced := []string{}
    block   := flag.Bool("block",   false, "Block domain of coins")
    update  := flag.Bool("update",  false, "Update list of coin domain")
    clean   := flag.Bool("clean",   false, "Unblock coin domain")

    flag.Parse()
    
    if *block  { choiced = append(choiced, "block" ) }
    if *update { choiced = append(choiced, "update") }
    if *clean  { choiced = append(choiced, "clean" ) }

    if len(choiced) == 1 {
        return choiced[0]
    }

    fmt.Println("Sorry, please choose only one option")
    flag.PrintDefaults()
    return "" 
}

// Replace miscellaneous IP Address
// with private ip address
func safeHostList(host string, re *regexp.Regexp) string {
    final   := re.ReplaceAllString(host, " ")
    domain  := strings.Fields(final)

    if len(domain) > 1 {
        domain[0]   = "0.0.0.0"
    }

    return strings.Join(domain, " ")
}


// Collector will collect all pointed host in os hosts file
func collector(hosts []string) (ori_list []string, coin_list []string) {
    bof := false
    re := regexp.MustCompile(`\s`)

    for _, host := range hosts {
        cmp_str := strings.ToLower(host)
        has_bof := strings.HasPrefix(cmp_str, "# bof d coins") 
        has_eof := strings.HasPrefix(cmp_str, "# eof d coins")

        if has_bof { 
            bof = true 
            continue
        }

        if has_eof { 
            bof = false 
            continue
        }


        if bof { 
            host      = safeHostList(host, re)
            coin_list = append(coin_list, host)   
        } else { 
            ori_list  = append(ori_list, host)    
        }
        
    }
    return ori_list, coin_list
}


// Collect all domain from current hosts
// then TutupBotol will get new hosts list from adblock list
// then merge it with current list of domain in our Host file
func updateList(config Config, hosts []string) {
    response, e := http.Get(config.UpdateEndpoint)
    err(e)
    defer response.Body.Close()

    content, e  := ioutil.ReadAll(response.Body)
    err(e)

    ori_list, _ := collector(hosts)

    s_ori_list  := strings.Join(ori_list,"\n")
    s_ori_list   = strings.Trim(s_ori_list, "\n")
    new_hosts   := []string{s_ori_list,"\n" , config.BofHeader, string(content), config.EofFooter}
    s_new_hosts := strings.Join(new_hosts, "\n")
    s_new_hosts  = strings.Trim(s_new_hosts, "\n")

    b_ori_host  := []byte(s_ori_list)
    ori_file_name := strings.Join([]string{config.HostFile, ".ori"}, "")
    ioutil.WriteFile(ori_file_name, b_ori_host, 0644)

    b_hosts := []byte(s_new_hosts)
    ioutil.WriteFile(config.HostFile, b_hosts, 0644)


}

// Clean List will remove all coin domains
// from our host file
func cleanList(config Config, hosts []string) {
    ori_list, _ := collector(hosts)

    s_ori_list  := strings.Join(ori_list,"\n")
    b_ori_host  := []byte(s_ori_list)

    ioutil.WriteFile(config.HostFile, b_ori_host, 0644)
}