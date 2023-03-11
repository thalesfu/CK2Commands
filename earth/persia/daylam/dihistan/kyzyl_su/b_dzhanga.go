package kyzyl_su

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 占加DzhangaBarony struct {
	feud.BaseBarony
}

var BDzhanga占加 feud.Barony = &占加DzhangaBarony{}

func init() {
    f := BDzhanga占加.(*占加DzhangaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "dzhanga",
		TitleName: "占加",
		TitleCode: "b_dzhanga",
	}
}
