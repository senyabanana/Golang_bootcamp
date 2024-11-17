package client

import (
	"context"
	"log"
	"math"
	"time"

	pb "anomaly-detection/internal/transmitter"
)

func AnomalyDetector(client pb.TransmitterServiceClient, k float64) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute*10)
	defer cancel()

	stream, err := client.StreamFrequencies(ctx, &pb.FrequencyRequest{})
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}

	var count int
	var sum, sumOfSquares float64

	for {
		msg, err := stream.Recv()
		if err != nil {
			log.Fatalf("error receiving message: %v", err)
		}

		frequency := msg.Frequency
		count++
		sum += frequency
		sumOfSquares += frequency * frequency

		mean := sum / float64(count)
		stddev := math.Sqrt((sumOfSquares / float64(count)) - (mean * mean))

		log.Printf("processed %d values. Mean: %.2f, STD: %.2f\n", count, mean, stddev)

		if math.Abs(frequency-mean) > k*stddev {
			log.Printf("anomaly detected: %.2f (Mean: %.2f, STD: %.2f)", frequency, mean, stddev)
		}
	}
}
