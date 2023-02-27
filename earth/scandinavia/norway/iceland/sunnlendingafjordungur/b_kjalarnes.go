package sunnlendingafjordungur

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 恰拉尔内斯KjalarnesBarony struct {
	feud.BaseBarony
}

var BKjalarnes恰拉尔内斯 feud.Barony = &恰拉尔内斯KjalarnesBarony{}

func init() {
    f := BKjalarnes恰拉尔内斯.(*恰拉尔内斯KjalarnesBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kjalarnes",
		TitleName: "恰拉尔内斯",
		TitleCode: "b_kjalarnes",
	}
}
