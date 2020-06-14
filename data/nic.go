package data

import (
	"context"
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

type NICData struct {
	Name   string
	Gid    int64
	Nid    int64
	Driver int32
}

func AddNIC(c *cobra.Command, data NICData) {
	base := etc.GetData(c)
	conn, err := grpc.Dial(base.Host, grpc.WithInsecure(), grpc.WithBlock(), grpc.WithTimeout(2*time.Second))
	if err != nil {
		log.Fatalf("Not connect; %v", err)
	}
	defer conn.Close()

	client := pb.NewNodeClient(conn)
	header := metadata.New(map[string]string{"authorization": base.Token})
	ctx := metadata.NewOutgoingContext(context.Background(), header)

	r, err := client.AddNIC(ctx, &pb.NICData{Name: data.Name, GroupID: data.Gid, NetID: data.Nid, Driver: data.Driver})
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Status: ")
	log.Println(r.Status)

	log.Printf("Info: ")
	log.Println(r.Info)
}

func DeleteNIC(c *cobra.Command, args []string) {

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

	r, err := client.DeleteNIC(ctx, &pb.NICData{ID: int64(id)})
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Status: ")
	log.Println(r.Status)

	log.Printf("Info: ")
	log.Println(r.Info)
}

func GetAllNIC(c *cobra.Command, args []string) {
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

	stream, err := client.GetAllNIC(ctx, &pb.Null{})
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
		tmp := []string{strconv.Itoa(int(d.ID)), d.Name, strconv.Itoa(int(d.GroupID)), strconv.Itoa(int(d.NetID)), driver, d.MacAddress, strconv.FormatBool(d.Lock)}
		data = append(data, tmp)
	}
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"ID", "Name", "GID", "NID", "Driver", "MAC", "LOCK"})

	for _, v := range data {
		table.Append(v)
	}
	table.Render()
}
