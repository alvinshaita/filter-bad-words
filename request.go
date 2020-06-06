package main

import (
	"fmt"
	"strconv"
	"os"
	"github.com/valyala/fasthttp"
)

var PORT = "8080"
var strRequestURI = []byte("http://localhost:"+PORT)

func post(client_request string){
	var strPost = []byte("POST")
	
	req := fasthttp.AcquireRequest()
	req.SetBody([]byte(client_request))
	req.Header.SetMethodBytes(strPost)
	req.SetRequestURIBytes(strRequestURI)
	req.Header.SetContentType("application/json")
	res := fasthttp.AcquireResponse()
	
	if err := fasthttp.Do(req, res); err != nil {
		panic("handle error")
	}
	
	fasthttp.ReleaseRequest(req)
	fasthttp.ReleaseResponse(res) // Only when you are done with body!
}

func get(){
	var strPost = []byte("GET")
	req := fasthttp.AcquireRequest()
	req.Header.SetMethodBytes(strPost)
	req.SetRequestURIBytes(strRequestURI)
	req.Header.SetContentType("application/json")
	res := fasthttp.AcquireResponse()
	
	if err := fasthttp.Do(req, res); err != nil {
		panic("handle error")
	}
	fasthttp.ReleaseRequest(req)
	body := res.Body()
	fasthttp.ReleaseResponse(res)
	fmt.Println(string(body))
}

func main(){
	if len(os.Args)==2{
		PORT = os.Args[1]
		strRequestURI = []byte("http://localhost:"+PORT)
	}
	
	fmt.Println("Sending request to port:", PORT)
	var i=0
	for i < 1000000 {
		post(`{"text": "test `+strconv.Itoa(i)+`"}`)
		get()
		i+=1
	}
}
