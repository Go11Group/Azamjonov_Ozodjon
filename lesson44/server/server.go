package main

import (
	"context"
	"flag"
	"fmt"
	pb "github.com/Azamjonov_Ozodjon/lesson44/proto/translator"
	"log"
	"net"

	"google.golang.org/grpc"
)

type TranslatorServer struct {
	EnglishWords map[string]string
	pb.UnimplementedTranslatorServerServer
}

func (s *TranslatorServer) Translate(ctx context.Context, in *pb.Request) (*pb.Response, error) {
	res := []string{}
	for _, word := range in.Words {
		if translatedWord, ok := s.EnglishWords[word]; ok {
			res = append(res, translatedWord)
		} else {
			return nil, fmt.Errorf("%s word not found in dictionary", word)
		}
	}
	return &pb.Response{Words: res}, nil
}

func main() {
	var translation = map[string]string{
		"apple":     "olma",
		"book":      "kitob",
		"house":     "uy",
		"car":       "mashina",
		"tree":      "daraxt",
		"water":     "suv",
		"fire":      "olov",
		"sky":       "osmon",
		"earth":     "yer",
		"sun":       "quyosh",
		"moon":      "oy",
		"star":      "yulduz",
		"flower":    "gul",
		"mountain":  "tog'",
		"river":     "daryo",
		"lake":      "ko'l",
		"forest":    "o'rmon",
		"animal":    "hayvon",
		"bird":      "qush",
		"fish":      "baliq",
		"bread":     "non",
		"milk":      "sut",
		"cheese":    "pishloq",
		"butter":    "yog'",
		"egg":       "tuxum",
		"meat":      "go'sht",
		"vegetable": "sabzavot",
		"fruit":     "meva",
		"school":    "maktab",
		"teacher":   "o'qituvchi",
		"student":   "talaba",
		"pen":       "ruchka",
		"pencil":    "qalam",
		"notebook":  "daftar",
		"computer":  "kompyuter",
		"phone":     "telefon",
		"window":    "deraza",
		"door":      "eshik",
		"chair":     "stul",
		"table":     "stol",
		"bed":       "karavot",
		"cup":       "fincan",
		"plate":     "likopcha",
		"spoon":     "qoshiTranslatorq",
		"fork":      "sanchqi",
		"knife":     "pichoq",
		"clothes":   "kiyim",
		"shoe":      "oyoq kiyim",
		"hat":       "shapka",
		"glasses":   "ko'zoynak",
		"watch":     "soat",
		"bag":       "sumka",
		"road":      "yo'l",
		"street":    "ko'cha",
		"city":      "shahar",
		"village":   "qishloq",
		"country":   "mamlakat",
		"world":     "dunyo",
		"family":    "oila",
		"friend":    "do'st",
		"love":      "sevgi",
		"happiness": "baxt",
		"sadness":   "g'am",
		"anger":     "g'azab",
		"fear":      "qo'rquv",
		"surprise":  "hayrat",
		"health":    "salomatlik",
		"medicine":  "dori",
		"doctor":    "shifokor",
		"hospital":  "kasalxona",
		"work":      "ish",
		"money":     "pul",
		"time":      "vaqt",
		"day":       "kun",
		"night":     "tun",
		"morning":   "ertalab",
		"afternoon": "tushdan keyin",
		"evening":   "kechqurun",
		"week":      "hafta",
		"month":     "oy",
		"year":      "yil",
		"spring":    "bahor",
		"summer":    "yoz",
		"autumn":    "kuz",
		"winter":    "qish",
	}
	flag.Parse()
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatal(err)
	}
	s := grpc.NewServer()
	pb.RegisterTranslatorServerServer(s, &TranslatorServer{EnglishWords: translation})
	log.Printf("Server is listening at %v", listener.Addr())
	err = s.Serve(listener)
	if err != nil {
		log.Fatal(err)
	}
}
