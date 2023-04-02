package main

import(
  "crypto/tls"
  "log"
)

func main() {
  conf := &tls.Config{InsecureSkipVerify: true}

  conn, err := tls.Dial("tcp", ":12947", conf)
  if (err != nil) {
    log.Fatal(err)
  }
  defer conn.Close()
  conn.Write([]byte("hello\n"))
}
