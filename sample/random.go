package sample

import (
	"grpc/pb"
	"math/rand"

	"github.com/google/uuid"
)

func randomKeyboardLayout() pb.Keyboard_Layout {
	switch rand.Intn(3) {
	case 0:
		return pb.Keyboard_QWERTY
	case 1:
		return pb.Keyboard_QWERTZ
	default:
		return pb.Keyboard_AZERTY
	}
}

func randomCPUBrand() string {
	return randomStringFromSet("Intel", "AMD")
}

func randomCPUName(brand string) string {
	if brand == "Intel" {
		return randomStringFromSet(
			"ultra 7", "ultra 5", "ultra 9",
		)
	} else {
		return randomStringFromSet(
			"ryzen 3",
			"ryzen 5",
			"ryzen 7",
			"ryzen 9",
		)
	}
}

func randomLaptopBrand() string {
	return randomStringFromSet("Apple", "Dell", "Lenovo")
}

func randomLaptopName(brand string) string {
	switch brand {
	case "Apple":
		return randomStringFromSet("M1", "M2", "M3 MacBook Pro")
	case "Dell":
		return randomStringFromSet("Latitude", "Vostro", "XPS", "Inspiron")
	default:
		return randomStringFromSet("Thinkpad X1", "ThinkPad P1", "Thinkpad P53")
	}
}

func randomGPUBrand() string {
	return randomStringFromSet("Nvidia", "AMD")
}

func randomGPUName(brand string) string {
	if brand == "Nvidia" {
		return randomStringFromSet(
			"RTX 3070",
			"RTX 4070",
			"RTX 5070 Ti",
		)
	} else {
		return randomStringFromSet(
			"RX 9070",
			"RX 9070XT",
			"RX 9060",
		)
	}
}

func randomScreenResolution() *pb.Screen_Resolution {
	height := randomInt(1080, 4320)
	width := height * 16 / 9
	resolution := &pb.Screen_Resolution{
		Height: uint32(height),
		Width:  uint32(width),
	}
	return resolution
}

func randomScreenPanel() pb.Screen_Panel {
	if rand.Intn(2) == 1 {
		return pb.Screen_IPS
	}
	return pb.Screen_OLED
}

func randomStringFromSet(a ...string) string {
	n := len(a)
	if n == 0 {
		return ""
	}
	return a[rand.Intn(n)]
}

func randomBool() bool {
	return rand.Intn(2) == 1
}

func randomInt(min int, max int) int {
	return min + rand.Intn(max-min+1)
}

func randomFloat64(min float64, max float64) float64 {
	return min + rand.Float64()*(max-min)
}

func randomFloat32(min float32, max float32) float32 {
	return min + rand.Float32()*(max-min)
}

func randomId() string {
	return uuid.New().String()
}
