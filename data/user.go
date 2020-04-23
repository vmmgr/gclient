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

type AuthData struct {
	Name  string
	Pass  string
	Token string
}

func AddUser(c *cobra.Command, args []string) {
	base := etc.GetData(c)
	conn, err := grpc.Dial(base.Host, grpc.WithInsecure(), grpc.WithBlock(), grpc.WithTimeout(time.Second))
	if err != nil {
		log.Fatalf("Not connect; %v", err)
	}
	defer conn.Close()

	client := pb.NewControllerClient(conn)
	header := metadata.New(map[string]string{"authorization": base.Token})
	ctx := metadata.NewIncomingContext(context.Background(), header)

	r, err := client.AddUser(ctx, &pb.UserData{Name: args[0], Pass: args[1]})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Status: ")
	log.Println(r.Status)

	log.Printf("Info: ")
	log.Println(r.Info)
}

func DeleteUser(c *cobra.Command, args []string) {
	base := etc.GetData(c)
	conn, err := grpc.Dial(base.Host, grpc.WithInsecure(), grpc.WithBlock(), grpc.WithTimeout(time.Second))
	if err != nil {
		log.Fatalf("Not connect; %v", err)
	}
	defer conn.Close()

	client := pb.NewControllerClient(conn)
	header := metadata.New(map[string]string{"authorization": base.Token})
	ctx := metadata.NewIncomingContext(context.Background(), header)

	r, err := client.DeleteUser(ctx, &pb.UserData{Id: args[0]})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Status: ")
	log.Println(r.Status)

	log.Printf("Info: ")
	log.Println(r.Info)
}

func GetAllUser(c *cobra.Command, args []string) {
	base := etc.GetData(c)
	conn, err := grpc.Dial(base.Host, grpc.WithInsecure(), grpc.WithBlock(), grpc.WithTimeout(time.Second))
	if err != nil {
		log.Fatalf("Not connect; %v", err)
	}
	defer conn.Close()

	client := pb.NewControllerClient(conn)
	header := metadata.New(map[string]string{"authorization": base.Token})
	ctx := metadata.NewIncomingContext(context.Background(), header)

	stream, err := client.GetAllUser(ctx, &pb.Null{})
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
		tmp := []string{d.Id, d.Name, strconv.Itoa(int(d.Auth)), d.Admingroup, d.Usergroup}
		data = append(data, tmp)
	}
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"UserID", "User", "Auth", "AdminGroup", "UserGroup"})

	for _, v := range data {
		table.Append(v)
	}
	table.Render()
}

func UserNameChange(c *cobra.Command, args []string) {
	base := etc.GetData(c)
	conn, err := grpc.Dial(base.Host, grpc.WithInsecure(), grpc.WithBlock(), grpc.WithTimeout(time.Second))
	if err != nil {
		log.Fatalf("Not connect; %v", err)
	}
	defer conn.Close()

	client := pb.NewControllerClient(conn)
	header := metadata.New(map[string]string{"authorization": base.Token})
	ctx := metadata.NewIncomingContext(context.Background(), header)

	r, err := client.UpdateUser(ctx, &pb.UserData{Id: args[0], Name: args[1]})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Status: ")
	log.Println(r.Status)

	log.Printf("Info: ")
	log.Println(r.Info)
}

func UserPassChange(c *cobra.Command, args []string) {
	base := etc.GetData(c)
	conn, err := grpc.Dial(base.Host, grpc.WithInsecure(), grpc.WithBlock(), grpc.WithTimeout(time.Second))
	if err != nil {
		log.Fatalf("Not connect; %v", err)
	}
	defer conn.Close()

	client := pb.NewControllerClient(conn)
	header := metadata.New(map[string]string{"authorization": base.Token})
	ctx := metadata.NewIncomingContext(context.Background(), header)

	r, err := client.UpdateUser(ctx, &pb.UserData{Id: args[0], Pass: args[1]})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Status: ")
	log.Println(r.Status)

	log.Printf("Info: ")
	log.Println(r.Info)
}
