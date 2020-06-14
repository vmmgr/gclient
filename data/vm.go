package data

import (
	"context"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/vmmgr/gclient/etc"
	pb "github.com/vmmgr/node/proto/proto-go"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"io"
	"log"
	"strconv"
	"time"
)

func AddVM(c *cobra.Command, args []string) {

	base := etc.GetData(c)
	conn, err := grpc.Dial(base.Host, grpc.WithInsecure(), grpc.WithBlock(), grpc.WithTimeout(2*time.Second))
	if err != nil {
		log.Fatalf("Not connect; %v", err)
	}
	defer conn.Close()

	client := pb.NewNodeClient(conn)
	header := metadata.New(map[string]string{"authorization": base.Token})
	ctx := metadata.NewOutgoingContext(context.Background(), header)
	groupID, _ := strconv.Atoi(args[1])
	driver, _ := strconv.Atoi(args[2])
	size, _ := strconv.Atoi(args[3])
	mode, _ := strconv.Atoi(args[4])

	stream, err := client.AddStorage(ctx, &pb.StorageData{
		Name: args[0], GroupID: int64(groupID), Driver: int32(driver), MaxSize: int64(size),
		Mode: int32(mode), Path: args[5], Image: args[6]})
	if err != nil {
		log.Fatal(err)
	}
	for {
		d, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Info: " + d.GetInfo())
		fmt.Println("Status: " + strconv.FormatBool(d.GetStatus()))
	}
}

func DeleteVM(c *cobra.Command, args []string) {

	base := etc.GetData(c)
	conn, err := grpc.Dial(base.Host, grpc.WithInsecure(), grpc.WithBlock(), grpc.WithTimeout(2*time.Second))
	if err != nil {
		log.Fatalf("Not connect; %v", err)
	}
	defer conn.Close()

	client := pb.NewNodeClient(conn)
	header := metadata.New(map[string]string{"authorization": base.Token})
	ctx := metadata.NewOutgoingContext(context.Background(), header)
	VMID, _ := strconv.Atoi(args[0])

	r, err := client.DeleteVM(ctx, &pb.VMData{
		ID: int64(VMID)})
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Status: ")
	log.Println(r.Status)

	log.Printf("Info: ")
	log.Println(r.Info)
}

func StatusVM(c *cobra.Command, args []string, status int32) {

	base := etc.GetData(c)
	conn, err := grpc.Dial(base.Host, grpc.WithInsecure(), grpc.WithBlock(), grpc.WithTimeout(2*time.Second))
	if err != nil {
		log.Fatalf("Not connect; %v", err)
	}
	defer conn.Close()

	client := pb.NewNodeClient(conn)
	header := metadata.New(map[string]string{"authorization": base.Token})
	ctx := metadata.NewOutgoingContext(context.Background(), header)
	vmID, _ := strconv.Atoi(args[0])

	r, err := client.UpdateVM(ctx, &pb.VMData{ID: int64(vmID), Mode: 0, Status: status})
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Status: ")
	log.Println(r.Status)

	log.Printf("Info: ")
	log.Println(r.Info)
}
