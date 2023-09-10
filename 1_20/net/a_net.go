package net

import (
	"log"
	"net"
)

func NetDemo() {
	//查看端口是否占用
	if _, err := net.ResolveTCPAddr("tcp", ":8080"); err != nil {
		log.Fatalf("解析监听地址异常，%s", err)
	}

	//获取所有网卡地址
	if ints, err := net.Interfaces(); err == nil {
		// https://en.wikipedia.org/wiki/IPv6_address#General_allocation
		_, ipv6Unicast, _ := net.ParseCIDR("2000::/3")
		log.Println("interfaces ", ints, " len: ", len(ints))
		for index, value := range ints {
			log.Println(index, "=", value)
			addrs, err := value.Addrs()
			if err != nil {
				log.Fatalln("error: ", err)
			}
			for index, address := range addrs {
				log.Println("----", index, "=", address)
				if ipnet, ok := address.(*net.IPNet); ok && ipnet.IP.IsGlobalUnicast() {
					log.Println(ipnet.IP.String(), "-=---", ipnet.Mask.String())
					_, bits := ipnet.Mask.Size()
					//需匹配全局单播地址
					if bits == 128 && ipv6Unicast.Contains(ipnet.IP) {
						log.Println("ipv6")
					}
					if bits == 32 {
						log.Println("ipv4")
					}
				}
			}
		}
	} else {
		log.Fatalln("interfaces err %s", err)
	}
}
