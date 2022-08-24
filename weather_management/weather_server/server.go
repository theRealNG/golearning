package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net"
	"sync"

	"google.golang.org/grpc"

	pb "github.com/theRealNG/golearning/weather_management/weather"
	"github.com/theRealNG/golearning/weather_management/weather_server/citydb"
	"github.com/theRealNG/golearning/weather_management/weather_server/openweather"
)

const (
	port = ":50051"
)

type WeatherServer struct {
	pb.UnimplementedWeatherServer
	mu sync.Mutex
}

// TODO: Add callback creat a go routine to update the weather for the city into db
func (s *WeatherServer) AddCity(ctx context.Context, in *pb.City) (*pb.City, error) {
	log.Printf("Received: %v", in.GetName())
	db := citydb.CityDB{}
	defer db.CloseConn()
	city := citydb.City{Name: in.GetName(), Lat: in.GetLat(), Long: in.GetLong()}
	db.CreateCity(&city)
	log.Printf("Created with ID: %v", city.ID)
	return in, nil
}

func (s *WeatherServer) MultiLangWeather(city *pb.City, stream pb.Weather_MultiLangWeatherServer) error {
	db := citydb.CityDB{}
	defer db.CloseConn()
	cityInfo := citydb.City{Name: city.GetName()}
	db.FindCity(&cityInfo)
	languages := [2]string{"en", "es"}
	for _, lang := range languages {
		response := openweather.GetWeather(&cityInfo, lang)
		resp := pb.WeatherInfo{
			Description:    response.Weather[0].Description,
			MaxTemperature: float32(response.Main.TempMax),
			MinTemperature: float32(response.Main.TempMin),
			FeelsLike:      float32(response.Main.FeelsLike),
		}
		if err := stream.Send(&resp); err != nil {
			return err
		}
	}
	return nil
}

func (s *WeatherServer) WeatherInLang(stream pb.Weather_WeatherInLangServer) error {
	var resp pb.WeatherInfoArray
	db := citydb.CityDB{}
	defer db.CloseConn()
	for {
		city, err := stream.Recv()
		if err == io.EOF {
			fmt.Println("Response to client streaming")
			fmt.Print(resp)
			return stream.SendAndClose(&resp)
		}
		if err != nil {
			return err
		}
		cityInfo := citydb.City{Name: city.GetName()}
		db.FindCity(&cityInfo)
		response := openweather.GetWeather(&cityInfo, "en")
		resp.WeatherInfos = append(resp.WeatherInfos, &pb.WeatherInfo{
			Description:    response.Weather[0].Description,
			MaxTemperature: float32(response.Main.TempMax),
			MinTemperature: float32(response.Main.TempMin),
			FeelsLike:      float32(response.Main.FeelsLike),
		})
	}
}

func (s *WeatherServer) MultiCityWeather(stream pb.Weather_MultiCityWeatherServer) error {
	db := citydb.CityDB{}
	defer db.CloseConn()
	for {
		in, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}

		cityInfo := citydb.City{Name: in.GetName()}
		db.FindCity(&cityInfo)
		for _, lang := range []string{"en", "es"} {
			response := openweather.GetWeather(&cityInfo, lang)
			resp := pb.WeatherInfo{
				Description:    response.Weather[0].Description,
				MaxTemperature: float32(response.Main.TempMax),
				MinTemperature: float32(response.Main.TempMin),
				FeelsLike:      float32(response.Main.FeelsLike),
			}
			if err := stream.Send(&resp); err != nil {
				return err
			}
		}
	}
}

func main() {
	db := citydb.CityDB{}
	db.SetupDB()
	db.CloseConn()

	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Failed to listed: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterWeatherServer(s, &WeatherServer{})
	log.Printf("Server listening at %v", lis.Addr())

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to server: %v", err)
	}
}
