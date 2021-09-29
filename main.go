package main

import (
    "fmt"
    "os"
    "bufio"
    "strings"
    "regexp"
  )

const SKIP_LINE = "__SKIP_LINE__"


func keyValueParser(line []string)(string, string){
    b_key := false
    var key string
    var value string
    len := len(line)
    // 「#」で始まる行はスキップ（コメント）
    if regexp.MustCompile(`^#`).Match([]byte(line[0])) {
        key = SKIP_LINE
        return key, value
    }
    for i := 0; i < len; i++ {
        if line[i] == "" { 
            continue 
        } else if !b_key {
            key = line[i]
            b_key = true
        }else{
            value = line[i]
        }
    }
    
    // keyに何も入らなければスキップ
    if !b_key{
        key = SKIP_LINE
    }
    return key, value
}


func main(){
    homeDir, _ := os.UserHomeDir()
    data, _ :=  os.Open(homeDir + "/.ssh/config")
    defer data.Close()
    scanner := bufio.NewScanner(data)
    for scanner.Scan(){
        line := strings.Split(scanner.Text()," ")
        key, value := keyValueParser(line)

        switch key {
        case "Host":
            fmt.Print("\n" + key + " " + value)
            break
        case "HostName":
            fmt.Print(" " + key + " " + value)
            break
        case SKIP_LINE:
            break
        default:
            break
        }
    }
}

