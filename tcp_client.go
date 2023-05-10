package main

import (
	"fmt"
	"log"
	"math"
	"net"
	"strconv"
	"strings"
	"time"
)

func main() {
   conn, err := net.Dial("tcp", ":8000")
   if nil != err {
      log.Println(err)
   }

   go func() {
      data := make([]byte, 4096)

      for {
         n, err := conn.Read(data)
         if err != nil {
            log.Println(err)
            return
         }

         log.Println("Server send : " + string(data[:n]))
         time.Sleep(time.Duration(3) * time.Second)
      }
   }()

   var username string;
   fmt.Print("Enter username: ");
   fmt.Scanln(&username);
   userId := RobinFingerprint(username);

   for {
      var s string
      fmt.Scanln(&s)
      payload := strconv.Itoa(userId) + " " + s;
      conn.Write([]byte(payload))
      time.Sleep(time.Duration(3) * time.Second)
   }
}


func RobinFingerprint(username string) int{
	username = strings.ToLower(username);
	var sum float64 = 0;
	for i := 0; i < len(username); i++ {
		sum += float64((float64(username[i]) * math.Pow(16, float64(i))));
	}
	return int(math.Mod(sum, 64)) ;
 }