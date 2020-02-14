package main
import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)
func handler(w http.ResponseWriter, req *http.Request) {
	requestID := req.Header.Get("x-fc-request-id")
	fmt.Println(fmt.Sprintf("sealyun robot Invoke Start RequestId: %s", requestID))
	defer func() {
		fmt.Println(fmt.Sprintf("FC Invoke End RequestId: %s", requestID))
	}()
	// your logic
	b, err := ioutil.ReadAll(req.Body)
	if err != nil {
		panic(err)
	}
	info := fmt.Sprintf("method =  %+v;\nheaders = %+v;\nbody = %+v", req.Method, req.Header, string(b))
	w.Write([]byte(fmt.Sprintf("Hello, golang  http invoke! detail:\n %s", info)))
}
func main() {
	fmt.Println("FunctionCompute go runtime inited.")
	http.HandleFunc("/", handler)
	port := os.Getenv("FC_SERVER_PORT")
	if port == "" {
		port = "9000"
	}
	http.ListenAndServe(":" + port, nil)
}