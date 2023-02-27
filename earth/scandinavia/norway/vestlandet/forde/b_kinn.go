package forde

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 欣KinnBarony struct {
	feud.BaseBarony
}

var BKinn欣 feud.Barony = &欣KinnBarony{}

func init() {
    f := BKinn欣.(*欣KinnBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kinn",
		TitleName: "欣",
		TitleCode: "b_kinn",
	}
}
