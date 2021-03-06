package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"time"
)

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func renewKerberosTicket() {
	ticker := time.NewTicker(1 * time.Hour)
	for t := range ticker.C {
		out, err := exec.Command("kinit", "-kt", "/etc/security/keytabs/hdfs-user.keytab", "hdfs-user").Output()
		checkErr(err)
		log.Printf("Kerberos Ticket for hdfs User Requested at: %v with output: %v\n", t, out)
	}
}

func main() {

	go renewKerberosTicket()

	log.Printf("Some work goes on here")
	
	select{}

}
