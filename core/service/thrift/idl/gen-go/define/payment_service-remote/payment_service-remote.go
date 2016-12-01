// Autogenerated by Thrift Compiler (0.9.3)
// DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING

package main

import (
	"define"
	"flag"
	"fmt"
	"git.apache.org/thrift.git/lib/go/thrift"
	"math"
	"net"
	"net/url"
	"os"
	"strconv"
	"strings"
)

func Usage() {
	fmt.Fprintln(os.Stderr, "Usage of ", os.Args[0], " [-h host:port] [-u url] [-f[ramed]] function [arg1 [arg2...]]:")
	flag.PrintDefaults()
	fmt.Fprintln(os.Stderr, "\nFunctions:")
	fmt.Fprintln(os.Stderr, "  Result CreatePaymentOrder(PaymentOrder o)")
	fmt.Fprintln(os.Stderr, "  PaymentOrder GetPaymentOrder(string paymentNo)")
	fmt.Fprintln(os.Stderr, "  PaymentOrder GetPaymentOrderById(i32 id)")
	fmt.Fprintln(os.Stderr, "  Result AdjustOrder(string paymentNo, double amount)")
	fmt.Fprintln(os.Stderr, "  Result DiscountByBalance(i32 orderId, string remark)")
	fmt.Fprintln(os.Stderr, "  DResult DiscountByIntegral(i32 orderId, i32 integral, bool ignoreOut)")
	fmt.Fprintln(os.Stderr, "  Result PaymentByPresent(i32 orderId, string remark)")
	fmt.Fprintln(os.Stderr, "  Result FinishPayment(string tradeNo, string spName, string outerNo)")
	fmt.Fprintln(os.Stderr)
	os.Exit(0)
}

func main() {
	flag.Usage = Usage
	var host string
	var port int
	var protocol string
	var urlString string
	var framed bool
	var useHttp bool
	var parsedUrl url.URL
	var trans thrift.TTransport
	_ = strconv.Atoi
	_ = math.Abs
	flag.Usage = Usage
	flag.StringVar(&host, "h", "localhost", "Specify host and port")
	flag.IntVar(&port, "p", 9090, "Specify port")
	flag.StringVar(&protocol, "P", "binary", "Specify the protocol (binary, compact, simplejson, json)")
	flag.StringVar(&urlString, "u", "", "Specify the url")
	flag.BoolVar(&framed, "framed", false, "Use framed transport")
	flag.BoolVar(&useHttp, "http", false, "Use http")
	flag.Parse()

	if len(urlString) > 0 {
		parsedUrl, err := url.Parse(urlString)
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error parsing URL: ", err)
			flag.Usage()
		}
		host = parsedUrl.Host
		useHttp = len(parsedUrl.Scheme) <= 0 || parsedUrl.Scheme == "http"
	} else if useHttp {
		_, err := url.Parse(fmt.Sprint("http://", host, ":", port))
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error parsing URL: ", err)
			flag.Usage()
		}
	}

	cmd := flag.Arg(0)
	var err error
	if useHttp {
		trans, err = thrift.NewTHttpClient(parsedUrl.String())
	} else {
		portStr := fmt.Sprint(port)
		if strings.Contains(host, ":") {
			host, portStr, err = net.SplitHostPort(host)
			if err != nil {
				fmt.Fprintln(os.Stderr, "error with host:", err)
				os.Exit(1)
			}
		}
		trans, err = thrift.NewTSocket(net.JoinHostPort(host, portStr))
		if err != nil {
			fmt.Fprintln(os.Stderr, "error resolving address:", err)
			os.Exit(1)
		}
		if framed {
			trans = thrift.NewTFramedTransport(trans)
		}
	}
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error creating transport", err)
		os.Exit(1)
	}
	defer trans.Close()
	var protocolFactory thrift.TProtocolFactory
	switch protocol {
	case "compact":
		protocolFactory = thrift.NewTCompactProtocolFactory()
		break
	case "simplejson":
		protocolFactory = thrift.NewTSimpleJSONProtocolFactory()
		break
	case "json":
		protocolFactory = thrift.NewTJSONProtocolFactory()
		break
	case "binary", "":
		protocolFactory = thrift.NewTBinaryProtocolFactoryDefault()
		break
	default:
		fmt.Fprintln(os.Stderr, "Invalid protocol specified: ", protocol)
		Usage()
		os.Exit(1)
	}
	client := define.NewPaymentServiceClientFactory(trans, protocolFactory)
	if err := trans.Open(); err != nil {
		fmt.Fprintln(os.Stderr, "Error opening socket to ", host, ":", port, " ", err)
		os.Exit(1)
	}

	switch cmd {
	case "CreatePaymentOrder":
		if flag.NArg()-1 != 1 {
			fmt.Fprintln(os.Stderr, "CreatePaymentOrder requires 1 args")
			flag.Usage()
		}
		arg97 := flag.Arg(1)
		mbTrans98 := thrift.NewTMemoryBufferLen(len(arg97))
		defer mbTrans98.Close()
		_, err99 := mbTrans98.WriteString(arg97)
		if err99 != nil {
			Usage()
			return
		}
		factory100 := thrift.NewTSimpleJSONProtocolFactory()
		jsProt101 := factory100.GetProtocol(mbTrans98)
		argvalue0 := define.NewPaymentOrder()
		err102 := argvalue0.Read(jsProt101)
		if err102 != nil {
			Usage()
			return
		}
		value0 := argvalue0
		fmt.Print(client.CreatePaymentOrder(value0))
		fmt.Print("\n")
		break
	case "GetPaymentOrder":
		if flag.NArg()-1 != 1 {
			fmt.Fprintln(os.Stderr, "GetPaymentOrder requires 1 args")
			flag.Usage()
		}
		argvalue0 := flag.Arg(1)
		value0 := argvalue0
		fmt.Print(client.GetPaymentOrder(value0))
		fmt.Print("\n")
		break
	case "GetPaymentOrderById":
		if flag.NArg()-1 != 1 {
			fmt.Fprintln(os.Stderr, "GetPaymentOrderById requires 1 args")
			flag.Usage()
		}
		tmp0, err104 := (strconv.Atoi(flag.Arg(1)))
		if err104 != nil {
			Usage()
			return
		}
		argvalue0 := int32(tmp0)
		value0 := argvalue0
		fmt.Print(client.GetPaymentOrderById(value0))
		fmt.Print("\n")
		break
	case "AdjustOrder":
		if flag.NArg()-1 != 2 {
			fmt.Fprintln(os.Stderr, "AdjustOrder requires 2 args")
			flag.Usage()
		}
		argvalue0 := flag.Arg(1)
		value0 := argvalue0
		argvalue1, err106 := (strconv.ParseFloat(flag.Arg(2), 64))
		if err106 != nil {
			Usage()
			return
		}
		value1 := argvalue1
		fmt.Print(client.AdjustOrder(value0, value1))
		fmt.Print("\n")
		break
	case "DiscountByBalance":
		if flag.NArg()-1 != 2 {
			fmt.Fprintln(os.Stderr, "DiscountByBalance requires 2 args")
			flag.Usage()
		}
		tmp0, err107 := (strconv.Atoi(flag.Arg(1)))
		if err107 != nil {
			Usage()
			return
		}
		argvalue0 := int32(tmp0)
		value0 := argvalue0
		argvalue1 := flag.Arg(2)
		value1 := argvalue1
		fmt.Print(client.DiscountByBalance(value0, value1))
		fmt.Print("\n")
		break
	case "DiscountByIntegral":
		if flag.NArg()-1 != 3 {
			fmt.Fprintln(os.Stderr, "DiscountByIntegral requires 3 args")
			flag.Usage()
		}
		tmp0, err109 := (strconv.Atoi(flag.Arg(1)))
		if err109 != nil {
			Usage()
			return
		}
		argvalue0 := int32(tmp0)
		value0 := argvalue0
		tmp1, err110 := (strconv.Atoi(flag.Arg(2)))
		if err110 != nil {
			Usage()
			return
		}
		argvalue1 := int32(tmp1)
		value1 := argvalue1
		argvalue2 := flag.Arg(3) == "true"
		value2 := argvalue2
		fmt.Print(client.DiscountByIntegral(value0, value1, value2))
		fmt.Print("\n")
		break
	case "PaymentByPresent":
		if flag.NArg()-1 != 2 {
			fmt.Fprintln(os.Stderr, "PaymentByPresent requires 2 args")
			flag.Usage()
		}
		tmp0, err112 := (strconv.Atoi(flag.Arg(1)))
		if err112 != nil {
			Usage()
			return
		}
		argvalue0 := int32(tmp0)
		value0 := argvalue0
		argvalue1 := flag.Arg(2)
		value1 := argvalue1
		fmt.Print(client.PaymentByPresent(value0, value1))
		fmt.Print("\n")
		break
	case "FinishPayment":
		if flag.NArg()-1 != 3 {
			fmt.Fprintln(os.Stderr, "FinishPayment requires 3 args")
			flag.Usage()
		}
		argvalue0 := flag.Arg(1)
		value0 := argvalue0
		argvalue1 := flag.Arg(2)
		value1 := argvalue1
		argvalue2 := flag.Arg(3)
		value2 := argvalue2
		fmt.Print(client.FinishPayment(value0, value1, value2))
		fmt.Print("\n")
		break
	case "":
		Usage()
		break
	default:
		fmt.Fprintln(os.Stderr, "Invalid function ", cmd)
	}
}
