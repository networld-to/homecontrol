package main

import (
	"fmt"
	"os"
	"strconv"
	"sync"
	"time"

	log "github.com/sirupsen/logrus"

	"golang.org/x/net/context"

	"github.com/networld-to/homecontrol/hue"
	"github.com/networld-to/homecontrol/version"
	"google.golang.org/grpc"
)

const (
	address = "127.0.0.1:50051"
)

// WithClientInterceptor : EXPERIMENTAL API
func WithClientInterceptor() grpc.DialOption {
	return grpc.WithUnaryInterceptor(clientInterceptor)
}

func clientInterceptor(
	ctx context.Context,
	method string,
	req interface{},
	reply interface{},
	cc *grpc.ClientConn,
	invoker grpc.UnaryInvoker,
	opts ...grpc.CallOption,
) error {
	start := time.Now()
	err := invoker(ctx, method, req, reply, cc, opts...) // <==
	log.Printf("invoke remote method=%s duration=%s error=%v", method,
		time.Since(start), err)
	return err
}

func must(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

func getGroups(client hue.LightsClient) *hue.Groups {
	resp, err := client.GetGroups(context.Background(), &hue.LightsRequest{})
	must(err)
	return resp
}

func switchOn(client hue.LightsClient, group int) {
	_, err := client.SwitchOn(context.Background(), &hue.LightsRequest{Group: int32(group), BrightnessPercent: 0.10})
	must(err)
	// log.Printf("Lights switched on: %v", r.Success)
}

func switchOff(client hue.LightsClient, group int) {
	_, err := client.SwitchOff(context.Background(), &hue.LightsRequest{Group: int32(group)})
	must(err)
	// log.Printf("Lights switched off: %v", r.Success)
}

func blink(client hue.LightsClient, group int) {
	_, err := client.Blink(context.Background(), &hue.LightsRequest{Group: int32(group), BrightnessPercent: 0.33})
	must(err)
	// log.Printf("Lights switched on: %v", r.Success)
}

func main() {
	start := time.Now()
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	must(err)
	defer conn.Close()
	c := hue.NewLightsClient(conn)

	cmd := ""
	if len(os.Args) > 1 {
		cmd = os.Args[1]
	}
	group := 2
	if len(os.Args) > 2 && os.Args[2] != "" {
		group, err = strconv.Atoi(os.Args[2])
		must(err)
	}
	switch {
	case cmd == "on":
		log.Printf("Switching on group %v", group)
		switchOn(c, group)
	case cmd == "off":
		log.Printf("Switching off group %v", group)
		switchOff(c, group)
	case cmd == "blink":
		log.Printf("Blink group %v", group)
		blink(c, group)
	case cmd == "perf":
		wg := &sync.WaitGroup{}
		n := 10
		wg.Add(n)
		for i := 0; i < n; i++ {
			go func(i int) {
				connection, _ := grpc.Dial(address, grpc.WithInsecure(), WithClientInterceptor())
				client := hue.NewLightsClient(connection)
				client.Echo(context.Background(), &hue.EchoMessage{Msg: "Hello, world " + fmt.Sprint(i)})
				// getGroups(c)
				wg.Done()
			}(i)
		}
		wg.Wait()
	case cmd == "version":
		log.WithField("version", version.Version).WithField("build", version.Build).Info("Version Information")
	default:
		resp := getGroups(c)
		log.Printf("Hue Light Groups: %v", resp.Groups)
	}

	end := time.Now()
	log.WithField("exec_time", end.Sub(start)).Printf("Execution finished")

}
