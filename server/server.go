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
	pb.UnsafeImageProcessorServer
}

func (s *Server) ConvertImage(ctx context.Context, req *pb.ConvertImageRequest) (*pb.ConvertImageResponse, error) {
	// Unmarshal image data from request (assuming it's a JPEG)
	img, err := decodeImage(req.ImageData)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "Invalid image data: %v", err)
	}

	// Convert to grayscale using a worker
	processedImg := convertImageToGrayscale(img)

	// Marshal processed image data back to byte slice
	processedData, err := encodeImage(processedImg)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Error processing image: %v", err)
	}

	return &pb.ConvertImageResponse{ProcessedData: processedData}, nil
}

func NewServer() *Server {
	return &Server{}
}
