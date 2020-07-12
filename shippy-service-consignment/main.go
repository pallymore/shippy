package main

import (
	"context"
	"log"

	// Import the generated protobuf code
	"github.com/micro/go-micro/v2"
	pb "github.com/pallymore/shippy/shippy-service-consignment/proto/consignment"
)

const port = ":50051"

type repository interface {
	Create(*pb.Consignment) (*pb.Consignment, error)
	GetAll() []*pb.Consignment
}

// Repository - Dummy repository, this simulates the use of a datastore
// of some kind. We'll replace this with a real implementation later on.
type Repository struct {
	consignments []*pb.Consignment
}

func (repo *Repository) Create(consignment *pb.Consignment) (*pb.Consignment, error) {
	updated := append(repo.consignments, consignment)
	repo.consignments = updated
	return consignment, nil
}

func (repo *Repository) GetAll() []*pb.Consignment {
	return repo.consignments
}

type consignmentService struct {
	repo repository
}

func (s *consignmentService) CreateConsignment(ctx context.Context, req *pb.Consignment, res *pb.Response) error {
	// Save our consignment
	consignment, err := s.repo.Create(req)
	if err != nil {
		return err
	}
	// Return matching the `Response` message we created in our
	// protobuf definition.
	res.Created = true
	res.Consignment = consignment
	return nil
}

func (s *consignmentService) GetConsignments(ctx context.Context, req *pb.GetRequest, res *pb.Response) error {
	consignments := s.repo.GetAll()
	res.Consignments = consignments
	return nil
}

func main() {
	repo := &Repository{}

	service := micro.NewService(
		// This name must match the package name given in your protobuf definition
		micro.Name("shippy.service.consignment"),
	)

	service.Init()

	// Register service
	if err := pb.RegisterShippingServiceHandler(service.Server(), &consignmentService{repo}); err != nil {
		log.Panic(err)
	}

	// Run the server
	if err := service.Run(); err != nil {
		log.Panic(err)
	}
}
