package main

import (
	"fmt"

	"net"
	"time"

	"github.com/sigurn/crc16"
)

func main() {
	//
	//dataBuff := bytes.NewBuffer([]byte{})

	//ID Token, SIA-DCS or ADM-CID
	var id string = "SIA-DCS"

	fmt.Println("id:", id)

	//seq from 0001-9999, four ASCII characters
	var iSeq uint16 = 1
	var seq string = fmt.Sprintf("%04d", iSeq)
	fmt.Println("seq:", seq)

	//Rrcvr. receiver number consists of an ASCII "R", follow by 1-6 HEX ASCII digits.
	var Rrcvr string = "579BD"

	Rrcvr = "R" + Rrcvr

	fmt.Println("Rrcvr:", Rrcvr)

	//Lpref. account prefix consists of an ASCII "L", follow by 1-6 HEX ASCII disits.
	var LPref string = "L000000"
	fmt.Println("LPref:", LPref)

	//acct. account number consists of an ASCII "#", followed by 3-16 ASCII characters representing hexadecimal digits

	var acct string = "Test123"
	fmt.Println("acct:", acct)

	acct = "#" + acct

	//[data] all data is in ASCII characters and the bracket characters "[" and "]" are included.

	var data string = "[" + acct + "|" + "NF1234/NPAZone123" + "]"
	fmt.Println("data:", data)

	//timestamp. the format of the timestamp is:<_HH:MM:SS,MM-DD-YYYY>. the braces are not part of the transmitted message. But the dunerscore, colon, commaand hyphen characters are included
	var timestamp string = "_16:08:34,03-06-2024"
	fmt.Println("timestamp:", timestamp)

	var data4crc string = "\"" + id + "\"" + seq + Rrcvr + LPref + acct + data + timestamp
	fmt.Println("data4crc:", data4crc, "length:", len(data4crc))

	//length of the message
	var lLLL int = len(data4crc)

	//convert 0LLL to 3 hex digits(in ASCII)
	var sLLL string = fmt.Sprintf("0%03x", lLLL)
	fmt.Println("sLLL:", sLLL)
	//calculate crc

	table := crc16.MakeTable(crc16.CRC16_ARC)
	crc := crc16.Checksum([]byte(data4crc), table)
	c := fmt.Sprintf("%X", crc)
	fmt.Println("CRC16==", c)

	//complete SIA data
	var siaData string = "\r" + c + sLLL + data4crc + "\n"
	fmt.Println("siaData:", siaData)

	//send sia data to platform by listening mode
	var network string = "tcp"
	var ip string = "10.7.66.16"
	var port string = "15646"
	var address string = ip + ":" + port
	conn, err := net.DialTimeout(network, address, time.Duration(5)*time.Second)
	if err != nil {
		fmt.Println("DialTimeout err:", err)
		return
	}

	defer conn.Close()

	nSend, err := conn.Write([]byte(siaData))
	if err != nil {
		fmt.Println("Write err:", err)
		return
	} else {
		fmt.Println("Write success, write len:", nSend)
	}

	time.Sleep(time.Duration(20) * time.Second)

}
