package server

import (
	"log"
	"math/rand"
	"time"

	pb "all-together/internal/transmitter"

	"github.com/google/uuid"
)

type Server struct {
	pb.UnimplementedTransmitterServiceServer
}

func (s *Server) StreamFrequencies(req *pb.FrequencyRequest, stream pb.TransmitterService_StreamFrequenciesServer) error {
	sessionID := uuid.New().String()
	mean := rand.Float64()*20 - 10
	std := rand.Float64()*1.2 + 0.3

	log.Printf("new session: %s, Mean: %.2f, StdDev: %.2f\n", sessionID, mean, std)

	for {
		frequency := rand.NormFloat64()*std + mean
		msg := pb.FrequencyMessage{
			SessionId: sessionID,
			Frequency: frequency,
			Timestamp: time.Now().UTC().Format(time.RFC3339),
		}

		if err := stream.Send(&msg); err != nil {
			log.Printf("error sending message: %v", err)
			return err
		}

		time.Sleep(time.Second)
	}
}
