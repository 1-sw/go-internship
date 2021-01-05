package main

import (
    "os"
    "fmt"
    "log"
    "io/ioutil"
    "encoding/base64"
)

const(
    OUTPUT_FILE      string = "result.txt"
    OBFUSCATE_OPTION string = "obfuscate"
    HELP_OPTION      string = "help"
    SPECIAL_SYMBOL_SLASH_N byte = 10
)

func output(text string,fileName string) {
    file, err := os.OpenFile(fileName,os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644) //CTRL+C CTRL+V
    if(err == nil) {
        file.WriteString(text+"\n")
        file.Close()
    }
}

func encode64(content string) string {
    return base64.StdEncoding.EncodeToString([]byte(content))
}

func main() {
    option := string(os.Args[1])
    if(option == OBFUSCATE_OPTION) {
        fileToObfuscate     := string(os.Args[2])
        keyBase             := string(os.Args[3])
        keySpecial          := string(os.Args[4])
        keyForSlashN        := keyBase+keySpecial+keyBase
        contentFromFile,err := ioutil.ReadFile(fileToObfuscate)
        if(err != nil) {
            fmt.Println("E: cannot find or create a file :")
            log.Fatal(err)
        }
        contentLength := int(len(contentFromFile))
        for currentSymbol := 0; currentSymbol < contentLength; currentSymbol++ {
            if(string(contentFromFile[currentSymbol])[0] == SPECIAL_SYMBOL_SLASH_N) {
                output(keyForSlashN,OUTPUT_FILE)
            } else {
                currentSymbol64 := encode64(string(contentFromFile[currentSymbol]))
                obfuscatedSymbol := keyBase + keySpecial + currentSymbol64 + keyBase + keySpecial
                output(obfuscatedSymbol,OUTPUT_FILE)
            }
        }
    } else if(option == HELP_OPTION) {
        fmt.Println("\nhelp:")
        fmt.Println("usage  : ./main [option] [file to obfuscate] [key 1] [key 2]")
        fmt.Println("explain: ./main obfuscate path_to_file random_letters_and_numbers_1 random_letters_and_numbers_1")
        fmt.Println("example: ./main obfuscate file.txt 0w9eid0w9di2093d 2039di2093di029d209dj\n")
    }
}


