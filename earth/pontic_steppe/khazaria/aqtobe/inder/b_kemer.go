package inder

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 凯梅尔KemerBarony struct {
	feud.BaseBarony
}

var BKemer凯梅尔 feud.Barony = &凯梅尔KemerBarony{}

func init() {
    f := BKemer凯梅尔.(*凯梅尔KemerBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kemer",
		TitleName: "凯梅尔",
		TitleCode: "b_kemer",
	}
}
