package sample

import (
	"math/rand"
	"time"

	pb "github.com/grpc-golang/pcbook/pb/proto"
)

func init(){
	rand.Seed(time.Now().UnixNano())
}

func randomKeyboardLayout() pb.Keyboard_Layout{
	switch rand.Intn(3){
	case 1:
		return pb.Keyboard_QWERTY
	case 2:
		return pb.Keyboard_QWERTZ
	default:
		return pb.Keyboard_AZERTY
	}
}

func randomBool() bool {
	return rand.Intn(2) == 1;
}

func randomCPUBrand() string{
	return randomStringFromSet("Intel", "AMD")
}

func randomStringFromSet(a ...string) string{
	n := len(a)
	if n == 0{
		return ""
	}
	return a[rand.Intn(n)]
}

func randomCPUName(brand string) string {
	if brand == "Intel" {
		return randomStringFromSet(
			"Core i7-9700HQ",
			"Core i9-9980K",
		)
	}

	return randomStringFromSet(
		"Ryzen 7 PRO 2700U",
		"Ryzen 5 PRO 3500U",
	)
}

func randomFloat64(min float64, max float64) float64{
	return min + rand.Float64() * (max - min)
}