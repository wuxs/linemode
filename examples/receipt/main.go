package main

import (
	"fmt"
	"github.com/jacobsa/go-serial/serial"
	"github.com/wuxs/linemode/pkg/linemode"
	"time"
)

func main() {
	//s, _ := net.Dial("tcp", "192.168.1.2:9100")

	options := serial.OpenOptions{
		PortName:        "/dev/ttyS1",
		BaudRate:        9600,
		DataBits:        8,
		StopBits:        1,
		MinimumReadSize: 4,
	}
	s, err := serial.Open(options)
	//s, _ := os.Create("test.txt")  // write to file
	defer s.Close()
	star := linemode.NewStar(s)
	star.Reset().
		SpecifyCodePage(linemode.Utf8).
		SpecifyLineSpace(48).
		SpecifyAlignment(linemode.Center).
		SpecifyBold().
		Print("车辆通行费\n").
		CancelBold().
		SpecifyAlignment(linemode.Left).
		Print("\n").
		Print(fmt.Sprintf("          发票代码: %s\n", "1234567890")).
		Print(fmt.Sprintf("          发票代码: %s\n", "1234567890")).
		Print("\n").
		Print(fmt.Sprintf("    车类: %s\n", "客一型")).
		Print(fmt.Sprintf("    车号: %s\n", "京A12345")).
		Print(fmt.Sprintf("    入口: %s\n", "西出口")).
		Print(fmt.Sprintf("    出口: %s\n", "东出口口")).
		Print(fmt.Sprintf("    日期: " + time.Now().Format("2006-01-01 15:04") + "\n")).
		FeedPaperLines(5).
		CutFull()
	_, err = star.Flush()
	if err != nil {
		fmt.Println(err.Error())
	}
}
