package hail

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 盖扎莱AlghazalahBarony struct {
	feud.BaseBarony
}

var BAlghazalah盖扎莱 feud.Barony = &盖扎莱AlghazalahBarony{}

func init() {
    f := BAlghazalah盖扎莱.(*盖扎莱AlghazalahBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "alghazalah",
		TitleName: "盖扎莱",
		TitleCode: "b_alghazalah",
	}
}
