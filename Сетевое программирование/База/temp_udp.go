package main

import(
    "fmt"
    "net"
    "log"
    "os"
)

func main() {
    fmt.Println("Сетевое приложение загружается....")
    if len(os.Args) < 2 {
        log.Fatal("Gde argumenti tvar")
        os.Exit(14)
    }
//     addr := fmt.Sprintf(":%v", os.Args[1])
    conn, err := net.Dial("tcp", os.Args[1])
    if err != nil {
        log.Fatal(err)
    }
//     fmt.Println("Listening at port", addr)
    buffer := make([]byte, 0xffff)
    for {
        _, err := conn.Read(buffer)
        if err != nil {
            log.Fatal(err)
        }
        fmt.Print(string(buffer))
    }
}
