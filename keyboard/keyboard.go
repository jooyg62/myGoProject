package keyboard

import (
        "log"
        "bufio"
        "os"
        "strconv"
        "strings"
)

func GetFloat() (float64, error) {
        reader := bufio.NewReader(os.Stdin)
        input, err := reader.ReadString('\n')
        if err != nil {
                log.Fatal(err)
        }

        input = strings.TrimSpace(input)
        number, err := strconv.ParseFloat(input, 64)
        if err != nil {
                log.Fatal(err)
        }

        return number, nil
}