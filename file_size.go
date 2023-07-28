// 파일 사이즈를 출력하는 프로그램
package main

import (
        "fmt"
        "log"
        "os"
)

func main() {
        fileInfo, err := os.Stat("my.txt")
        if err != nil {
                log.Fatal(err)
        }
        fmt.Println(fileInfo)

}