package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"strconv"
	"strings"
	"time"

	"github.com/go-redis/redis/v8"
)

func main() {
	ctx := context.Background()

	// Create Redis client
	rdb := redis.NewClient(&redis.Options{
		Addr: "localhost:6379", // Redis server address
	})

	rand.Seed(time.Now().UnixNano())

	// Define the minimum and maximum range
	minimum := 1000
	maximum := 9999999

	var oldprice int
	var memory []int
	companyName := "ExampleCompany" // Replace with the actual company name

	for i := 0; i < 1000; i++ {
		randomNumber := rand.Intn(maximum-minimum+1) + minimum
		time.Sleep(time.Second)

		// Publish messages to different channels
		messagee := fmt.Sprintf("New price: %d", randomNumber)
		memory = append(memory, randomNumber)

		currentTime := time.Now().Format(time.RFC3339)

		minPriceKey := fmt.Sprintf("minPrice:%s", companyName)
		maxPriceKey := fmt.Sprintf("maxPrice:%s", companyName)

		// Find the current min and max in memory
		currentMin := min(memory)
		currentMax := max(memory)

		if randomNumber <= currentMin {
			minPriceValue := fmt.Sprintf("price: %d, time: %s", randomNumber, currentTime)
			err := rdb.Set(ctx, minPriceKey, minPriceValue, 0).Err()
			if err != nil {
				log.Fatalf("Failed to set minPrice: %v", err)
			} else {
				fmt.Printf("Updated Redis with new minPrice: %s\n", minPriceValue)
			}
		}

		if randomNumber >= currentMax {
			maxPriceValue := fmt.Sprintf("price: %d, time: %s", randomNumber, currentTime)
			err := rdb.Set(ctx, maxPriceKey, maxPriceValue, 0).Err()
			if err != nil {
				log.Fatalf("Failed to set maxPrice: %v", err)
			} else {
				fmt.Printf("Updated Redis with new maxPrice: %s\n", maxPriceValue)
			}
		}

		err := rdb.Publish(ctx, "stockPriceUpdater", messagee).Err()
		if err != nil {
			log.Fatalf("Failed to publish to stockPriceUpdater: %v", err)
		}

		memoryStr := joinInts(memory)
		err = rdb.Publish(ctx, "stockPriceMonitor", memoryStr).Err()
		if err != nil {
			log.Fatalf("Failed to publish to stockPriceMonitor: %v", err)
		}

		if i != 0 {
			var changeMessage string
			if randomNumber < oldprice {
				changeMessage = "decreased"
			} else if randomNumber > oldprice {
				changeMessage = "increased"
			} else {
				changeMessage = "no changes"
			}

			err = rdb.Publish(ctx, "stockAnalytics", changeMessage).Err()
			if err != nil {
				log.Fatalf("Failed to publish to stockAnalytics: %v", err)
			}
		}
		oldprice = randomNumber
		fmt.Println("Messages published successfully")
	}
}

func min(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	minNum := nums[0]
	for _, num := range nums {
		if num < minNum {
			minNum = num
		}
	}
	return minNum
}

func max(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	maxNum := nums[0]
	for _, num := range nums {
		if num > maxNum {
			maxNum = num
		}
	}
	return maxNum
}

func joinInts(nums []int) string {
	strs := make([]string, len(nums))
	for i, num := range nums {
		strs[i] = strconv.Itoa(num)
	}
	return strings.Join(strs, " ")
}
