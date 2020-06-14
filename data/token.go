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

func GenerateToken(c *cobra.Command, args []string, user, pass string) {
	base := etc.GetData(c)
	conn, err := grpc.Dial(base.Host, grpc.WithInsecure(), grpc.WithBlock(), grpc.WithTimeout(time.Second))
	if err != nil {
		log.Fatalf("Not connect; %v", err)
	}
	defer conn.Close()

	client := pb.NewControllerClient(conn)
	header := metadata.New(map[string]string{"authorization": base.Token})
	ctx := metadata.NewOutgoingContext(context.Background(), header)

	mode, err := strconv.Atoi(args[0])
	if err != nil {
		log.Fatalf("Error: string to int")
	}

	r, err := client.GenerateToken(ctx, &pb.UserData{Name: user, Pass: pass, Auth: int32(mode)})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	if r.Status {
		log.Println("ok")
		log.Printf("User: " + r.Data1 + "| Token: ")
		log.Println(r.Data2)
	} else {
		log.Println("ng")
		log.Printf("Info: ")
		log.Println(r.Info)
	}

}

func DeleteToken(c *cobra.Command, args []string) {
	base := etc.GetData(c)
	conn, err := grpc.Dial(base.Host, grpc.WithInsecure(), grpc.WithBlock(), grpc.WithTimeout(time.Second))
	if err != nil {
		log.Fatalf("Not connect; %v", err)
	}
	defer conn.Close()

	client := pb.NewControllerClient(conn)
	header := metadata.New(map[string]string{"authorization": base.Token})
	ctx := metadata.NewOutgoingContext(context.Background(), header)

	r, err := client.DeleteToken(ctx, &pb.Null{})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Status: ")
	log.Println(r.Status)

	log.Printf("Info: ")
	log.Println(r.Info)
}

func GetAllToken(c *cobra.Command, args []string) {
	base := etc.GetData(c)
	conn, err := grpc.Dial(base.Host, grpc.WithInsecure(), grpc.WithBlock(), grpc.WithTimeout(time.Second))
	if err != nil {
		log.Fatalf("Not connect; %v", err)
	}
	defer conn.Close()

	client := pb.NewControllerClient(conn)
	header := metadata.New(map[string]string{"authorization": base.Token})
	ctx := metadata.NewOutgoingContext(context.Background(), header)

	stream, err := client.GetAllToken(ctx, &pb.Null{})
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
		begintime := time.Unix(d.Begintime, 0).Format("2006/01/02 15:04:05")
		endtime := time.Unix(d.Endtime, 0).Format("2006/01/02 15:04:05")

		tmp := []string{strconv.Itoa(int(d.Id)), d.Token, strconv.Itoa(int(d.Userid)), begintime, endtime}
		data = append(data, tmp)
	}
	log.Println(time.Now().Format("2006/01/02 15:04:05"))
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"ID", "Token", "UserID", "BeginTime", "EndTime"})

	for _, v := range data {
		table.Append(v)
	}
	table.Render()
}
