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
	"strings"
	"time"
)

type NetData struct {
	Name string
	Gid  string
	Vlan int32
}

func AddNet(c *cobra.Command, data NetData) {
	var gid []int64

	base := etc.GetData(c)
	conn, err := grpc.Dial(base.Host, grpc.WithInsecure(), grpc.WithBlock(), grpc.WithTimeout(2*time.Second))
	if err != nil {
		log.Fatalf("Not connect; %v", err)
	}
	defer conn.Close()

	client := pb.NewNodeClient(conn)
	header := metadata.New(map[string]string{"authorization": base.Token})
	ctx := metadata.NewOutgoingContext(context.Background(), header)

	for _, a := range strings.Split(data.Gid, ",") {
		tmp, _ := strconv.Atoi(a)
		gid = append(gid, int64(tmp))
	}

	r, err := client.AddNet(ctx, &pb.NetData{Name: data.Name, GroupID: gid, VLAN: data.Vlan})
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Status: ")
	log.Println(r.Status)

	log.Printf("Info: ")
	log.Println(r.Info)
}

func DeleteNet(c *cobra.Command, args []string) {

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

	r, err := client.DeleteNet(ctx, &pb.NetData{ID: int64(id)})
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Status: ")
	log.Println(r.Status)

	log.Printf("Info: ")
	log.Println(r.Info)
}

func GetAllNet(c *cobra.Command, args []string) {
	base := etc.GetData(c)
	conn, err := grpc.Dial(base.Host, grpc.WithInsecure(), grpc.WithBlock(), grpc.WithTimeout(time.Second))
	if err != nil {
		log.Fatalf("Not connect; %v", err)
	}
	defer conn.Close()

	client := pb.NewNodeClient(conn)
	header := metadata.New(map[string]string{"authorization": base.Token})
	ctx := metadata.NewOutgoingContext(context.Background(), header)
	var gid string

	stream, err := client.GetAllNet(ctx, &pb.Null{})
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
		for _, d := range d.GroupID {
			gid += strconv.Itoa(int(d)) + ","
		}
		tmp := []string{strconv.Itoa(int(d.ID)), d.Name, gid, strconv.Itoa(int(d.VLAN)), strconv.FormatBool(d.Lock)}
		data = append(data, tmp)
	}
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"ID", "Name", "GID", "VLAN", "LOCK"})

	for _, v := range data {
		table.Append(v)
	}
	table.Render()
}
