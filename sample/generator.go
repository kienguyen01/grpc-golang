package sample

import (
	"math/rand"

	"github.com/golang/protobuf/ptypes"
	"github.com/google/uuid"
	pb "github.com/grpc-golang/pcbook/pb/proto"
)

func NewKeyBoard() *pb.Keyboard {
	keyboard := &pb.Keyboard{
		Layout:  randomKeyboardLayout(),
		Backlit: randomBool(),
	}

	return keyboard
}

func NewCPU() *pb.CPU {
	brand := randomCPUBrand()
	name := "Intel"

	numberCores := 2 + rand.Intn(8-2+1)
	numberThread := numberCores + rand.Intn(12-numberCores+1)

	minGhz := randomFloat64(2.0, 3.5)
	maxGhz := randomFloat64(minGhz, 5.0)

	cpu := &pb.CPU{
		Brand:   brand,
		Name:    name,
		Core:    uint32(numberCores),
		Threads: uint32(numberThread),
		MinGhz:  minGhz,
		MaxGhz:  maxGhz,
	}

	return cpu
}

func NewGPU() *pb.GPU {
	brand := randomStringFromSet("AMD", "NVIDIA")

	minGhz := randomFloat64(1.0, 1.5)
	maxGhz := randomFloat64(minGhz, 2.0)

	memory := &pb.Memory{
		Value: uint64(2 + rand.Intn(6-2+1)),
		Unit:  pb.Memory_GIGABYTE,
	}

	gpu := &pb.GPU{
		Brand:  brand,
		Name:   "GTX 1660Ti",
		MinGhz: minGhz,
		MaxGhz: maxGhz,
		Memory: memory,
	}

	return gpu
}

func NewRAM() *pb.Memory{
	ram := &pb.Memory{
		Value: uint64(4 + rand.Intn(64 - 4 + 1)),
		Unit:  pb.Memory_GIGABYTE,
	}

	return ram
}

func NewSSD() *pb.Storage {
	ssd := &pb.Storage{
		Driver: pb.Storage_SSD,
		Memory: &pb.Memory{
			Value: uint64(128 + rand.Intn(1024 - 128 + 1)),
			Unit:  pb.Memory_GIGABYTE,
		},
	}

	return ssd
}

func NewHDD() *pb.Storage{
	hdd := &pb.Storage{
		Driver: pb.Storage_HDD,
		Memory: &pb.Memory{
			Value: uint64(1 + rand.Intn(6 - 1 + 1)),
			Unit:  pb.Memory_TERABYTE,
		},
	}

	return hdd
}

func NewScreen() *pb.Monitor{
	height := 1080
	width := 1920
	screen := &pb.Monitor{
		Multitouch: false,
		SizeInch: 15,
		Resolution: &pb.Monitor_Resolution{
			Height: uint32(height),
			Width: uint32(width),
		},
		Panel: pb.Monitor_IPS,
	}

	return screen
}

func NewLaptop() *pb.Laptop{
	brand := "Razer"
	name := "Blade"
	laptop := &pb.Laptop{
		Id: uuid.New().String(),
		Brand: brand,
		Name: name,
		Cpu: NewCPU(),
		Ram: NewRAM(),
		Gpus: []*pb.GPU{NewGPU()},
		Storages: []*pb.Storage{NewSSD(), NewHDD()},
		Monitor: NewScreen(),
		Keyboard: NewKeyBoard(),
		Weight: &pb.Laptop_WeightKg{
			WeightKg: 2.5,
		},
		PriceUsd: 2000,
		ReleaseYear: 2019,
		UpdatedAt: ptypes.TimestampNow(),
	}

	return laptop
}

