package usecase

import (
	"math/rand"
	"time"
)

const (
	toshiaki = "俊"
	mika     = "美佳"
	yoko     = "お母さん"
	seiichi  = "パチ"
)

var familyMembers = []string{
	toshiaki,
	mika,
	yoko,
	seiichi,
}

func PickMember() string {
	rand.Seed(time.Now().UnixNano())
	idx := rand.Intn(len(familyMembers))
	return familyMembers[idx]
}
