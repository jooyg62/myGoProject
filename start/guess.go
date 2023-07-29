// guess 프로그램은 플레이어가 난수를 맞히는 게임입니다.
package main

import (
        "bufio"
        "fmt"
        "log"
        "math/rand"
        "os"
        "strconv"
        "strings"
        "time"
)

func main() {
        seconds := time.Now().Unix()
        rand.Seed(seconds)

        // 1부터 100 사이의 난수 생성
        randomNumber := rand.Intn(100) + 1

        // fmt.Println("생성된 난수:", randomNumber)

        var chances int
        for x := 0; x < 10; x++ {
                fmt.Print("숫자를 맞춰보세요: ")

                reader := bufio.NewReader(os.Stdin)
                input, err := reader.ReadString('\n')
                if err != nil {
                        log.Fatal(err)
                }

                input = strings.TrimSpace(input)
                guessNumber, err := strconv.Atoi(input)
                if err != nil {
                        log.Fatal(err)
                }

                var status string
                if randomNumber == guessNumber {
                        status = "EQUAL"
                } else if randomNumber > guessNumber {
                        status = "LOW"
                } else {
                        status = "HIGH"
                }

                if status == "EQUAL" {
                        fmt.Println("Good job! You guessed it!")
                        os.Exit(0)
                } else {
                        fmt.Println("Oops, Your Guess was ", status)
                }

                chances = 10 - (x + 1)
                fmt.Println("You have ", chances, " more chances.")
        }

        fmt.Println("You didn't guess my number. It was: randomNumber")
}