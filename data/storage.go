package data

import (
	"context"
	"fmt"
	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
	"github.com/vmmgr/gclient/etc"
	pb "github.com/vmmgr/node/proto/proto-go"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"io"
	"log"
	"os"
	"strconv"
	"time"
)

type StorageData struct {
	Name   string
	Gid    int64
	Driver int32
	Size   int64
	Mode   int32
	Path   string
	Image  string
}

func AddStorage(c *cobra.Command, data StorageData) {
	base := etc.GetData(c)
	conn, err := grpc.Dial(base.Host, grpc.WithInsecure(), grpc.WithBlock(), grpc.WithTimeout(2*time.Second))
	if err != nil {
		log.Fatalf("Not connect; %v", err)
	}
	defer conn.Close()

	client := pb.NewNodeClient(conn)
	header := metadata.New(map[string]string{"authorization": base.Token})
	ctx := metadata.NewOutgoingContext(context.Background(), header)

	stream, err := client.AddStorage(ctx, &pb.StorageData{
		Name: data.Name, GroupID: data.Gid, Driver: data.Driver, MaxSize: data.Size,
		Mode: data.Mode, Path: data.Path, Image: data.Image})
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

func DeleteStorage(c *cobra.Command, args []string) {

	base := etc.GetData(c)
	conn, err := grpc.Dial(base.Host, grpc.WithInsecure(), grpc.WithBlock(), grpc.WithTimeout(2*time.Second))
	if err != nil {
		log.Fatalf("Not connect; %v", err)
	}
	defer conn.Close()

	client := pb.NewNodeClient(conn)
	header := metadata.New(map[string]string{"authorization": base.Token})
	ctx := metadata.NewOutgoingContext(context.Background(), header)
	id, _ := strconv.Atoi(args[0])

	r, err := client.DeleteStorage(ctx, &pb.StorageData{
		ID: int64(id)})
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Status: ")
	log.Println(r.Status)

	log.Printf("Info: ")
	log.Println(r.Info)
}

func GetAllStorage(c *cobra.Command, args []string) {
	base := etc.GetData(c)
	conn, err := grpc.Dial(base.Host, grpc.WithInsecure(), grpc.WithBlock(), grpc.WithTimeout(time.Second))
	if err != nil {
		log.Fatalf("Not connect; %v", err)
	}
	defer conn.Close()

	client := pb.NewNodeClient(conn)
	header := metadata.New(map[string]string{"authorization": base.Token})
	ctx := metadata.NewOutgoingContext(context.Background(), header)
	var driver string

	stream, err := client.GetAllStorage(ctx, &pb.Null{})
	if err != nil {
		log.Fatal(err)
	}
	var data [][]string
	for {
		d, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		if d.Driver == 1 {
			driver = "virtio"
		}
		tmp := []string{strconv.Itoa(int(d.ID)), strconv.Itoa(int(d.GroupID)), d.Name, driver, d.Path, strconv.Itoa(int(d.MaxSize))}
		data = append(data, tmp)
	}
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"ID", "GID", "Name", "Driver", "Path", "MaxSize"})

	for _, v := range data {
		table.Append(v)
	}
	table.Render()
}
