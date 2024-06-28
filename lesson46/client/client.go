package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	pbTransport "github.com/Azamjonov_Ozodjon/lesson46/genproto/generator/transport"
	pbWeather "github.com/Azamjonov_Ozodjon/lesson46/genproto/generator/weather"

	"google.golang.org/grpc"
)

func main() {
	// Set default values for service and method
	service := "weather"          // Default service
	method := "GetCurrentWeather" // Default method

	if len(os.Args) >= 2 {
		service = os.Args[1]
	}
	if len(os.Args) >= 3 {
		method = os.Args[2]
	}

	switch service {
	case "weather":
		runWeatherServiceClient(method)
	case "transport":
		runTransportServiceClient(method)
	default:
		fmt.Println("Unknown service:", service)
	}
}

func runWeatherServiceClient(method string) {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pbWeather.NewWeatherServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	switch method {
	case "GetCurrentWeather":
		r, err := c.GetCurrentWeather(ctx, &pbWeather.CurrentWeatherRequest{Location: "DefaultLocation"})
		if err != nil {
			log.Fatalf("could not get current weather: %v", err)
		}
		fmt.Printf("Current Weather: Temperature: %.1f, Humidity: %.1f%%, Wind Speed: %.1f m/s\n", r.Temperature, r.Humidity, r.WindSpeed)
	case "GetWeatherForecast":
		r, err := c.GetWeatherForecast(ctx, &pbWeather.WeatherForecastRequest{Location: "DefaultLocation", Days: 2})
		if err != nil {
			log.Fatalf("could not get weather forecast: %v", err)
		}
		fmt.Println("Weather Forecast:")
		for _, forecast := range r.Forecasts {
			fmt.Printf("%s: Temperature: %.1f, Humidity: %.1f%%, Wind Speed: %.1f m/s\n", forecast.Date, forecast.Temperature, forecast.Humidity, forecast.WindSpeed)
		}
	case "ReportWeatherCondition":
		r, err := c.ReportWeatherCondition(ctx, &pbWeather.ReportWeatherConditionRequest{Location: "DefaultLocation", Condition: "Sunny"})
		if err != nil {
			log.Fatalf("could not report weather condition: %v", err)
		}
		fmt.Printf("Report Weather Condition: Success: %v\n", r.Success)
	default:
		fmt.Println("Unknown method:", method)
	}
}

func runTransportServiceClient(method string) {
	conn, err := grpc.Dial("localhost:50052", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pbTransport.NewTransportServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	switch method {
	case "GetBusSchedule":
		r, err := c.GetBusSchedule(ctx, &pbTransport.BusScheduleRequest{BusNumber: "42"})
		if err != nil {
			log.Fatalf("could not get bus schedule: %v", err)
		}
		fmt.Println("Bus Schedule:")
		for _, schedule := range r.Schedules {
			fmt.Printf("Time: %s, Destination: %s\n", schedule.Time, schedule.Destination)
		}
	case "TrackBusLocation":
		r, err := c.TrackBusLocation(ctx, &pbTransport.BusLocationRequest{BusNumber: "42"})
		if err != nil {
			log.Fatalf("could not track bus location: %v", err)
		}
		fmt.Printf("Bus Location: %s\n", r.Location)
	case "ReportTrafficJam":
		r, err := c.ReportTrafficJam(ctx, &pbTransport.TrafficJamReportRequest{Location: "Main Street", Description: "Heavy traffic"})
		if err != nil {
			log.Fatalf("could not report traffic jam: %v", err)
		}
		fmt.Printf("Report Traffic Jam: Success: %v\n", r.Success)
	default:
		fmt.Println("Unknown method:", method)
	}
}
