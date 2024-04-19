package server

import (
	"context"

	pb "github.com/menothe/ipg/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// service interface
type ImageProcessorServer interface {
	ConvertImage(ctx context.Context, req *pb.ConvertImageRequest) (*pb.ConvertImageResponse, error)
}

// Server implementation
type Server struct {
	WorkerPool chan func()
	pb.UnsafeImageProcessorServer
}

func (s *Server) ConvertImage(ctx context.Context, req *pb.ConvertImageRequest) (*pb.ConvertImageResponse, error) {
	// Unmarshal image data from request (assuming it's a JPEG)
	img, err := decodeImage(req.ImageData)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "Invalid image data: %v", err)
	}

	// Convert to grayscale using a worker
	processedImg := convertImageToGrayscale(img, s.WorkerPool)

	// Marshal processed image data back to byte slice
	processedData, err := encodeImage(processedImg)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Error processing image: %v", err)
	}

	return &pb.ConvertImageResponse{ProcessedData: processedData}, nil
}

func NewServer(poolSize int) *Server {
	s := &Server{
		WorkerPool: make(chan func(), poolSize),
	}
	// Pre-populate the worker pool with worker functions
	for i := 0; i < poolSize; i++ {
		worker := func() {
			for {
				task, ok := <-s.WorkerPool
				if !ok {
					break // Worker pool is closing
				}
				task()
			}
		}
		go worker()
	}
	return s
}
