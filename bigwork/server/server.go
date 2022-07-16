package server

import (
	pb "bigwork/gen/go/pet/v1"
	"context"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/genproto/googleapis/type/datetime"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"time"
)

type petServer struct {
	pets []pb.Pet
}

func NewHelloServer() pb.PetStoreServiceServer  {
	s := petServer{
	}
	return &s
}

func (p *petServer) PutPet(ctx context.Context, req *pb.PutPetRequest) (*pb.PutPetResponse, error) {
	//名字不能为空
	if req.Name == ""{
		return nil,status.Error(codes.InvalidArgument,"the pet Name can't be empty")
	}
	//类型要在定义的范围内
	if req.PetType < pb.PetType_PET_TYPE_CAT || req.PetType > pb.PetType_PET_TYPE_HAMSTER {
		return nil,status.Error(codes.InvalidArgument,"the pet type no exist")
	}

	tim := time.Now()
	zoneName, offset:= tim.Zone()
	pet :=pb.Pet{}
	pet.Name = req.Name
	pet.PetType = req.PetType
	pet.PetId = primitive.NewObjectID().Hex()
	pet.CreatedAt= &datetime.DateTime{
		Year: int32(tim.Year()),
		Month: int32(tim.Month()),
		Day: int32(tim.Day()),
		Hours: int32(tim.Hour()),
		Minutes: int32(tim.Minute()),
		Seconds: int32(tim.Second()),
		Nanos: int32(tim.UnixNano()),
		TimeOffset:&datetime.DateTime_TimeZone{
			TimeZone: &datetime.TimeZone{
				Id: string(offset),
				Version: zoneName,
			},
		},
	}
	p.pets = append(p.pets, pet)

	log.Info("PutPet Success")
	return &pb.PutPetResponse{Pet: &pet},nil
}


func (p *petServer) GetPet(ctx context.Context, req *pb.GetPetRequest) (*pb.GetPetResponse, error) {
	for _,pet := range p.pets{
		if pet.PetId == req.PetId{
			log.Info("GetPet Success")
			return &pb.GetPetResponse{Pet: &pet},nil
		}
	}

	return nil,status.Error(codes.InvalidArgument,"can't find this pet")
}

func (p *petServer) DeletePet(ctx context.Context, req *pb.DeletePetRequest) (*pb.DeletePetResponse, error) {
	for _,pet := range p.pets{
		if pet.PetId == req.PetId{
			log.Info("DeletePet Success")
			return &pb.DeletePetResponse{},nil
		}
	}

	return nil,status.Error(codes.InvalidArgument,"this pet no exist")
}
