package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"time"

	pb "github.com/theRealNG/golearning/weather_management/weather"
	"google.golang.org/grpc"
)

const (
	address = "localhost:50051"
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())

	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewWeatherClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	result, err := c.AddCity(ctx, &pb.City{Name: "Hyd", Lat: 17.3850, Long: 78.4867})
	if err != nil {
		log.Fatalf("Failed to create city: %v", err)
	}
	log.Printf("Recieved: %v", result.GetName())

	stream, err := c.MultiLangWeather(ctx, &pb.City{Name: "Hyd"})
	if err != nil {
		log.Fatalf("Failed to get information: %v", err)
	}
	for {
		info, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("call failed: %v", err)
		}
		fmt.Println(info)
	}

	fmt.Println("Streaming Client Example")
	streamCity, err := c.WeatherInLang(ctx)
	if err != nil {
		log.Fatalf("Streaming failed")
	}

	for _, city := range []pb.City{pb.City{Name: "Hyd"}, pb.City{Name: "Hyd"}} {
		if err := streamCity.Send(&city); err != nil {
			log.Fatalf("client.streamCity: streamCity.Send(%v) failed: %v", city, err)
		}
	}
	if reply, err := streamCity.CloseAndRecv(); err != io.EOF {
		fmt.Println(reply)
	}
	if err != nil {
		log.Fatalf("Error while recieving %v", err)
	}

	fmt.Println("BiDirectional Streaming Client Example")
	streamCityWeather, err := c.MultiCityWeather(ctx)
	if err != nil {
		log.Fatalf("Streaming failed")
	}
	waitc := make(chan struct{})
	go func() {
		for {
			in, err := streamCityWeather.Recv()
			if err == io.EOF {
				// read done.
				close(waitc)
				return
			}
			if err != nil {
				log.Fatalf("client.MultiCityWeather failed: %v", err)
			}
			log.Println(in)
		}
	}()
	for _, city := range []pb.City{pb.City{Name: "Hyd"}, pb.City{Name: "Hyd"}} {
		if err := streamCityWeather.Send(&city); err != nil {
			log.Fatalf("client.MultiCityWeather: streamCityWeather.Send(%v) failed: %v", city, err)
		}
	}
	streamCityWeather.CloseSend()
	<-waitc
}
