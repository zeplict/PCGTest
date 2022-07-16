package main

import (
	pb "bigwork/gen/go/pet/v1"
	"context"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"os"
	"strconv"
)

func main()  {
	conn, err := grpc.Dial("127.0.0.1:8000", grpc.WithInsecure())
	if err !=nil{
		log.Info("err")
	}
	c := pb.NewPetStoreServiceClient(conn)

	cmd := os.Args[1]
	if cmd == "put"{
		//装载参数
		req := pb.PutPetRequest{}
		t, _ := strconv.Atoi(os.Args[2]) //参数2为类型,参数3为名称
		req.PetType = pb.PetType(t)
		req.Name = os.Args[3]
		//请求
		putRes,err := c.PutPet(context.Background(),&req)
		if err !=nil{
			log.Error("PutPet fail :",err)
			return
		}
		log.WithFields(log.Fields{
			"id":putRes.Pet.PetId,
			"name":putRes.Pet.PetId,
			"type":putRes.Pet.PetType}).Info("PutPet success")

	}else if cmd == "get" {
		//装载参数
		req := pb.GetPetRequest{PetId: os.Args[2]}
		//请求
		getRes,err := c.GetPet(context.Background(),&req)
		if err !=nil{
			log.Error("GetPet:",err)
			return
		}
		log.WithFields(log.Fields{
			"id":getRes.Pet.PetId,
			"name":getRes.Pet.PetId,
			"type":getRes.Pet.PetType}).Info("PutPet success")

	}else if cmd == "delete" {
		//装载参数
		req := pb.DeletePetRequest{PetId: os.Args[2]}
		//请求
		_,err := c.DeletePet(context.Background(),&req)
		if err !=nil{
			log.Error("DeletePet:",err)
			return
		}
		log.Info("DeletePet success :")
	}




}