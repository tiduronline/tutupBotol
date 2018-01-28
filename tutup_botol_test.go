package main

import (
    "testing"
    "fmt"
    "strings"
    "regexp"
    
)

func assertEqual(t *testing.T, ret_val interface{}, comp_val interface{}, message string) {
    if ret_val == comp_val { return }

    if len(message) > 0 {
        message = fmt.Sprintf("%v != %v", ret_val, comp_val)
    }

    t.Fatal(message)
}

func TestSafeHostList(t *testing.T) {
    var tmp_safe_list string
    re := regexp.MustCompile(`\s`)
    
    domains     := []string{"0.0.0.0 domain1", "192.168.0.1 domain2", "8.8.8.8 google.com", "192.168.0.1\tdomain3"}
    safe_list   := []string{}

    for _, host := range domains {
        tmp_safe_list = safeHostList(host, re)
        safe_list = append(safe_list, tmp_safe_list)
    }

    s_safe_list := strings.Join(safe_list, " ")

    should_str := "0.0.0.0 domain1 0.0.0.0 domain2 0.0.0.0 google.com 0.0.0.0 domain3"
    assertEqual(t, s_safe_list, should_str , s_safe_list)
}


func TestCollector(t *testing.T) {
    list := []string{}
    list = append(list, "127.0.0.1 localhost")
    list = append(list, "127.0.0.1 development")
    list = append(list, "# BOF D COINS BLOCK STATUS")
    list = append(list, "0.0.0.0 coin1")
    list = append(list, "0.0.0.0 coin2")
    list = append(list, "0.0.0.0 coin3")
    list = append(list, "0.0.0.0 coin4")
    list = append(list, "# EOF D COINS")
    
    ori_list    := "127.0.0.1 localhost 127.0.0.1 development"
    coin_list   := "0.0.0.0 coin1 0.0.0.0 coin2 0.0.0.0 coin3 0.0.0.0 coin4"

    ori_data, coin_data := collector(list)
    s_ori   := strings.Join(ori_data, " ")
    s_coin  := strings.Join(coin_data, " ")

    assertEqual(t, s_ori, ori_list, "Ori List isnot match")
    assertEqual(t, s_coin, coin_list, "Coin List isnot match")
} 


// func TestMSlice(t *testing.T) {
//     str := m_slice("contoh string", 0, 6)
//     assertEqual(t, str, "contoh", "Slice failed")
// }