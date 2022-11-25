package main

import (
	"fmt"
)

//https://api.coingecko.com/api/v3/coins/list

/*примеры запросов:
см фото кода сегодня

*/

func main() {

/*	var nSec int
	fmt.Println("Enter timeout(sec):")
	fmt.Scanf("%d", &nSec)
	fmt.Println("Timeout:", nSec, "sec")

	requests := make([]string, 0, 0)
	reqtimes := make([]time.Time, 0, 0)
	
	timeout := time.Duration(nSec) * time.Second
		
	for {
		var req string
		fmt.Scanf("%s", &req)
		if len(req) == 0 {
			continue
		}
		if req == "q" {
			break
		}
		ctime := time.Now()	

		//finding new time interval start
		if len(reqtimes) == 0 {
		//do nothing here
		} else if ctime.Sub(reqtimes[len(reqtimes) - 1]) > timeout {
			requests = make([]string, 0, 0)
			reqtimes = make([]time.Time, 0, 0)
		} else { 
			for i := len(reqtimes) - 2; i >= 0; i--{
				if ctime.Sub(reqtimes[i]) > timeout {
					fmt.Println("Time remove", reqtimes[i])
					requests = requests[i+1:len(requests)]
					reqtimes = reqtimes[i+1:len(reqtimes)]
					break
				}
			}
		}

		found := false
		for i:=0; i < len(requests); i++{
			if req == requests[i] {
				fmt.Println("REJECTED:",req, ctime)
				found = true
				break
			}
		}
		if !found {
			requests = append(requests, req)
			reqtimes = append(reqtimes, ctime)
			fmt.Println("ACCEPTED:",req, ctime)
		}
	}
*/
	fmt.Println("Finished")
}
