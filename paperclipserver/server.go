package paperclipserver

import (
	"context"

	pb "github.com/zachgoldstein/twirpPaperclips/rpc"
)

// Server implements the UniversalPaperclips service
type Server struct {
	PaperClips int32
}

// NewServer creates an instance of our server
func NewServer() *Server {
	return &Server{
		PaperClips: 1,
	}
}

// GetPaperclips returns the current paperclip count
func (s *Server) GetPaperclips(ctx context.Context, empty *pb.Empty) (*pb.Paperclips, error) {
	return &pb.Paperclips{
		Paperclips: s.PaperClips,
	}, nil
}

// IncrementPaperclips increments the paperclip count
func (s *Server) IncrementPaperclips(ctx context.Context, size *pb.Size) (*pb.Empty, error) {
	s.PaperClips += size.Paperclips
	return &pb.Empty{}, nil
}

// CalculateUniverseLifespan calulcates the universe's lifespan and returns some marginally relaxing value
func (s *Server) CalculateUniverseLifespan(ctx context.Context, empty *pb.Empty) (*pb.Dread, error) {
	return &pb.Dread{
		Paperclips:       s.PaperClips,
		UniverseLifespan: "42",
	}, nil
}
