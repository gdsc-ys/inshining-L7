package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"net"
	"os"
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

         fmt.Println("Server send : " + string(data[:n]))
         time.Sleep(time.Duration(3) * time.Second)
      }
   }()

   var username string;
   fmt.Print("Enter username: ");
   fmt.Scanln(&username);
   userId := RobinFingerprint(username);

   for {
      var s string
      scanner := bufio.NewScanner(os.Stdin);
      scanner.Scan();
      s = scanner.Text();

      // Header
      contentLenght := len(s);
      t := time.Now();
      time_format := t.Format(time.RFC3339);
      header := strconv.Itoa(userId) + "Content-Length:" + strconv.Itoa(contentLenght) + "," + "language:en,"  + "date-time:" + time_format + "\n\n";

      payload :=  s;
      packet := header + payload;
      bin := StringToBin(packet);
      bin = "01" + bin
      conn.Write([]byte(bin))
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

func StringToBin(s string) string {
   res := ""
   for _, c := range s {
       res = fmt.Sprintf("%s%.8b", res, c)
   }
   return res
}