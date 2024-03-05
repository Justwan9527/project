package main

import (
	"fmt"

	"github.com/sigurn/crc16"
)

func main() {
	//
	//dataBuff := bytes.NewBuffer([]byte{})

	//length of the message
	var lLLL int = 64

	//convert 0LLL to 3 hex digits(in ASCII)
	var sLLL string = fmt.Sprintf("0%03x", lLLL)

	fmt.Println("sLLL:", sLLL)

	//ID Token, SIA-DCS or ADM-CID
	var id string = "ADM-CID"

	fmt.Println("id:", id)

	//seq from 0001-9999, four ASCII characters
	var iSeq uint16 = 1
	var seq string = fmt.Sprintf("%04d", iSeq)
	fmt.Println("seq:", seq)

	//Rrcvr. receiver number consists of an ASCII "R", follow by 1-6 HEX ASCII digits.
	var Rrcvr string

	fmt.Println("Rrcvr:", Rrcvr)

	//Lpref. account prefix consists of an ASCII "L", follow by 1-6 HEX ASCII disits.
	var LPref string = "L000000"
	fmt.Println("LPref:", LPref)

	//acct. account number consists of an ASCII "#", followed by 3-16 ASCII characters representing hexadecimal digits

	var acct string = "#1234"
	fmt.Println("acct:", acct)

	//[data] all data is in ASCII characters and the bracket characters "[" and "]" are included.

	var data string = "[" + acct + "|" + "1140 00 007" + "]"
	fmt.Println("data:", data)

	//timestamp. the format of the timestamp is:<_HH:MM:SS,MM-DD-YYYY>. the braces are not part of the transmitted message. But the dunerscore, colon, commaand hyphen characters are included
	var timestamp string = "_22:49:34,01-22-2012"
	fmt.Println("timestamp:", timestamp)

	var data4crc string = "\"" + id + "\"" + seq + Rrcvr + LPref + acct + data + timestamp
	fmt.Println("data4crc:", data4crc, "length:", len(data4crc))

	//calculate crc

	table := crc16.MakeTable(crc16.CRC16_ARC)
	crc := crc16.Checksum([]byte(data4crc), table)
	c := fmt.Sprintf("%x", crc)
	fmt.Println("CRC16==", c)

	//binary.Write(dataBuff)

}
