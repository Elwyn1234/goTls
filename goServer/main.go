package main

import(
  "crypto/tls"
  "bufio"
  "log"
)

func main() {
  cer, _ := tls.LoadX509KeyPair("server.crt", "server.key")
  config := &tls.Config{Certificates: []tls.Certificate{cer}}
  ln, err := tls.Listen("tcp", ":12947", config)
  if (err != nil) {
    log.Fatal(err)
  }
  defer ln.Close()
  
  for {
    println("Waiting for message")
    conn, _ := ln.Accept()
    println("Message received")
    r := bufio.NewReader(conn)
    for {
      msg, err := r.ReadString('\n')
      println(msg)
      if (err != nil) {break}
    }
    conn.Close()
  }
}
