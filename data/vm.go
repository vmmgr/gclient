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

type VmData struct {
	Gid     int64
	Name    string
	CPU     int32
	Mem     int32
	Storage string
	NIC     string
	PCI     string
}

func AddVM(c *cobra.Command, data VmData) {
	var storage []*pb.StorageData
	var nic []*pb.NICData
	var pci []*pb.PCIData

	base := etc.GetData(c)
	conn, err := grpc.Dial(base.Host, grpc.WithInsecure(), grpc.WithBlock(), grpc.WithTimeout(2*time.Second))
	if err != nil {
		log.Fatalf("Not connect; %v", err)
	}
	defer conn.Close()

	client := pb.NewNodeClient(conn)
	header := metadata.New(map[string]string{"authorization": base.Token})
	ctx := metadata.NewOutgoingContext(context.Background(), header)

	for _, tmp := range strings.Split(data.Storage, ",") {
		id, _ := strconv.Atoi(tmp)
		storage = append(storage, &pb.StorageData{ID: int64(id)})
	}
	for _, tmp := range strings.Split(data.NIC, ",") {
		id, _ := strconv.Atoi(tmp)
		nic = append(nic, &pb.NICData{ID: int64(id)})
	}
	for _, tmp := range strings.Split(data.PCI, ",") {
		id, _ := strconv.Atoi(tmp)
		pci = append(pci, &pb.PCIData{ID: int64(id)})
	}

	r, err := client.AddVM(ctx, &pb.VMData{
		Name: data.Name, GroupID: data.Gid, CPU: data.CPU, Mem: data.Mem, Storage: storage, NIC: nic, PCIData: pci})
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Status: ")
	log.Println(r.Status)

	log.Printf("Info: ")
	log.Println(r.Info)
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

func GetAllVM(c *cobra.Command, args []string) {
	base := etc.GetData(c)
	conn, err := grpc.Dial(base.Host, grpc.WithInsecure(), grpc.WithBlock(), grpc.WithTimeout(time.Second))
	if err != nil {
		log.Fatalf("Not connect; %v", err)
	}
	defer conn.Close()

	client := pb.NewNodeClient(conn)
	header := metadata.New(map[string]string{"authorization": base.Token})
	ctx := metadata.NewOutgoingContext(context.Background(), header)
	var storage string
	var nic string
	var pci string

	stream, err := client.GetAllVM(ctx, &pb.Null{})
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
		for _, a := range d.Storage {
			storage += strconv.Itoa(int(a.ID)) + ","
		}
		for _, a := range d.NIC {
			nic += strconv.Itoa(int(a.ID)) + ","
		}
		for _, a := range d.PCIData {
			pci += strconv.Itoa(int(a.ID)) + ","
		}

		tmp := []string{strconv.Itoa(int(d.ID)), strconv.Itoa(int(d.GroupID)), d.Name, strconv.Itoa(int(d.CPU)), strconv.Itoa(int(d.Mem)), storage, nic, pci}
		data = append(data, tmp)
	}
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"ID", "GroupID", "Name", "CPU", "Mem", "Storage", "NIC", "PCI"})

	for _, v := range data {
		table.Append(v)
	}
	table.Render()
}
