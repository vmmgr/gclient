package data

import (
	"context"
	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
	pb "github.com/vmmgr/controller/proto/proto-go"
	"github.com/vmmgr/gclient/etc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"io"
	"log"
	"os"
	"strconv"
	"time"
)

type MaxSpec struct {
	MaxVM      int32
	MaxCPU     int32
	MaxMem     int32
	MaxStorage int64
}

func AddGroup(c *cobra.Command, args []string, spec MaxSpec) {
	base := etc.GetData(c)
	conn, err := grpc.Dial(base.Host, grpc.WithInsecure(), grpc.WithBlock(), grpc.WithTimeout(time.Second))
	if err != nil {
		log.Fatalf("Not connect; %v", err)
	}
	defer conn.Close()

	mode, err := strconv.Atoi(args[1])
	if err != nil {
		log.Fatal("Error: string to int")
	}

	client := pb.NewControllerClient(conn)
	header := metadata.New(map[string]string{"authorization": base.Token})
	ctx := metadata.NewOutgoingContext(context.Background(), header)

	r, err := client.AddGroup(ctx, &pb.GroupData{
		Name: args[0],
		Mode: int32(mode),
		Sepc: &pb.SpecData{Vm: spec.MaxVM, Cpu: spec.MaxCPU, Mem: spec.MaxMem, Storage: spec.MaxStorage},
	})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Status: ")
	log.Println(r.Status)

	log.Printf("Info: ")
	log.Println(r.Info)
}

func DeleteGroup(c *cobra.Command, args []string) {
	base := etc.GetData(c)
	conn, err := grpc.Dial(base.Host, grpc.WithInsecure(), grpc.WithBlock(), grpc.WithTimeout(time.Second))
	if err != nil {
		log.Fatalf("Not connect; %v", err)
	}
	defer conn.Close()

	client := pb.NewControllerClient(conn)
	header := metadata.New(map[string]string{"authorization": base.Token})
	ctx := metadata.NewOutgoingContext(context.Background(), header)

	r, err := client.DeleteGroup(ctx, &pb.GroupData{Id: args[0]})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Status: ")
	log.Println(r.Status)

	log.Printf("Info: ")
	log.Println(r.Info)
}

func GetAllGroup(c *cobra.Command, args []string) {
	base := etc.GetData(c)
	conn, err := grpc.Dial(base.Host, grpc.WithInsecure(), grpc.WithBlock(), grpc.WithTimeout(time.Second))
	if err != nil {
		log.Fatalf("Not connect; %v", err)
	}
	defer conn.Close()

	client := pb.NewControllerClient(conn)
	header := metadata.New(map[string]string{"authorization": base.Token})
	ctx := metadata.NewOutgoingContext(context.Background(), header)

	stream, err := client.GetAllGroup(ctx, &pb.Null{})
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
		tmp := []string{d.Id, d.Name, d.Admin, d.User, strconv.Itoa(int(d.Sepc.Vm)), strconv.Itoa(int(d.Sepc.Cpu)),
			strconv.Itoa(int(d.Sepc.Mem)), strconv.Itoa(int(d.Sepc.Storage)), d.Sepc.Net}
		data = append(data, tmp)
	}
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"ID", "Name", "Admin", "User", "MaxVM", "MaxCPU", "MaxMem", "MaxStorage", "Net"})

	for _, v := range data {
		table.Append(v)
	}
	table.Render()
}

func GetGroup(c *cobra.Command, args []string) {
	base := etc.GetData(c)
	conn, err := grpc.Dial(base.Host, grpc.WithInsecure(), grpc.WithBlock(), grpc.WithTimeout(time.Second))
	if err != nil {
		log.Fatalf("Not connect; %v", err)
	}
	defer conn.Close()

	client := pb.NewControllerClient(conn)
	header := metadata.New(map[string]string{"authorization": base.Token})
	ctx := metadata.NewOutgoingContext(context.Background(), header)

	stream, err := client.GetGroup(ctx, &pb.GroupData{Id: args[0]})
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
		tmp := []string{d.Id, d.Name, d.Admin, d.User, strconv.Itoa(int(d.Sepc.Vm)), strconv.Itoa(int(d.Sepc.Cpu)),
			strconv.Itoa(int(d.Sepc.Mem)), strconv.Itoa(int(d.Sepc.Storage)), d.Sepc.Net}
		data = append(data, tmp)
	}
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"ID", "Name", "Admin", "User", "MaxVM", "MaxCPU", "MaxMem", "MaxStorage", "Net"})

	for _, v := range data {
		table.Append(v)
	}
	table.Render()
}

func JoinAddGroup(c *cobra.Command, args []string) {
	base := etc.GetData(c)
	d := pb.GroupData{Id: args[1], Mode: 0}
	conn, err := grpc.Dial(base.Host, grpc.WithInsecure(), grpc.WithBlock(), grpc.WithTimeout(time.Second))
	if err != nil {
		log.Fatalf("Not connect; %v", err)
	}
	defer conn.Close()

	client := pb.NewControllerClient(conn)
	header := metadata.New(map[string]string{"authorization": base.Token})
	ctx := metadata.NewOutgoingContext(context.Background(), header)

	if args[0] == "0" {
		//Admin
		d.Admin = args[2]
	} else if args[0] == "1" {
		//User
		d.User = args[2]
	}
	r, err := client.JoinGroup(ctx, &d)
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Status: ")
	log.Println(r.Status)

	log.Printf("Info: ")
	log.Println(r.Info)
}

func JoinDeleteGroup(c *cobra.Command, args []string) {
	base := etc.GetData(c)
	d := pb.GroupData{Id: args[1], Mode: 1}
	conn, err := grpc.Dial(base.Host, grpc.WithInsecure(), grpc.WithBlock(), grpc.WithTimeout(time.Second))
	if err != nil {
		log.Fatalf("Not connect; %v", err)
	}
	defer conn.Close()

	client := pb.NewControllerClient(conn)
	header := metadata.New(map[string]string{"authorization": base.Token})
	ctx := metadata.NewOutgoingContext(context.Background(), header)

	if args[0] == "0" {
		//Admin
		d.Admin = args[2]
	} else if args[0] == "1" {
		//User
		d.User = args[2]
	}
	r, err := client.JoinGroup(ctx, &d)
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Status: ")
	log.Println(r.Status)

	log.Printf("Info: ")
	log.Println(r.Info)
}

func GroupNameChange(c *cobra.Command, args []string) {
	base := etc.GetData(c)
	conn, err := grpc.Dial(base.Host, grpc.WithInsecure(), grpc.WithBlock(), grpc.WithTimeout(time.Second))
	if err != nil {
		log.Fatalf("Not connect; %v", err)
	}
	defer conn.Close()

	client := pb.NewControllerClient(conn)
	header := metadata.New(map[string]string{"authorization": base.Token})
	ctx := metadata.NewOutgoingContext(context.Background(), header)

	r, err := client.UpdateGroup(ctx, &pb.GroupData{Id: args[0], Name: args[1]})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Status: ")
	log.Println(r.Status)

	log.Printf("Info: ")
	log.Println(r.Info)
}

func GroupSpecChange(c *cobra.Command, args []string, spec MaxSpec) {
	base := etc.GetData(c)
	conn, err := grpc.Dial(base.Host, grpc.WithInsecure(), grpc.WithBlock(), grpc.WithTimeout(time.Second))
	if err != nil {
		log.Fatalf("Not connect; %v", err)
	}
	defer conn.Close()

	client := pb.NewControllerClient(conn)
	header := metadata.New(map[string]string{"authorization": base.Token})
	ctx := metadata.NewOutgoingContext(context.Background(), header)

	r, err := client.UpdateGroup(ctx, &pb.GroupData{
		Id:   args[0],
		Sepc: &pb.SpecData{Vm: spec.MaxVM, Cpu: spec.MaxCPU, Mem: spec.MaxMem, Storage: spec.MaxStorage},
	})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Status: ")
	log.Println(r.Status)

	log.Printf("Info: ")
	log.Println(r.Info)
}
