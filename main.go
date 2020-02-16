package main

import (

	// grp1 "end-kfgrpc/protos"
	// "time"

	"github.com/dabory/abango"
	// "google.golang.org/grpc"
	// "google.golang.org/grpc/keepalive"
)

func init() {
	abango.RunEndRequest() //Should be here before gRpcServer Run
}

func main() {

}

// func gRpcRequest(askname string, askstr string) (string, string) {

// 	dial, err := grpc.Dial(
// 		cf.XConfig["gRpcAddr"]+":"+cf.XConfig["gRpcPort"],
// 		grpc.WithInsecure(),
// 		grpc.WithKeepaliveParams(keepalive.ClientParameters{
// 			Time:                time.Millisecond,
// 			Timeout:             time.Millisecond,
// 			PermitWithoutStream: true,
// 		}),
// 	)

// 	if err != nil {
// 		e.MyErr("WOEIURQPWERQ", nil, true)
// 	}
// 	defer dial.Close()
// 	c := grp1.NewGrp1Client(dial)

// 	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
// 	defer cancel()

// 	r, err := c.StdRpc(ctx, &grp1.StdAsk{AskName: askname, AskStr: askstr})
// 	if err != nil {
// 		e.MyErr("erttryrtye", err, true)
// 	}

// 	return r.RetStr, r.RetSta

// }
