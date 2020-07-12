package main

import (
	pb "github.com/pallymore/shippy/shippy/service-vessel/proto/vessel"
)

type Repository interface {
	FindAvailable(*pb.Specification) (*pb.Vessel, error)
}
