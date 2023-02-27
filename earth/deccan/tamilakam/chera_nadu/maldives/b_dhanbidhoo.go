package maldives

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 达恩比杜DhanbidhooBarony struct {
	feud.BaseBarony
}

var BDhanbidhoo达恩比杜 feud.Barony = &达恩比杜DhanbidhooBarony{}

func init() {
    f := BDhanbidhoo达恩比杜.(*达恩比杜DhanbidhooBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "dhanbidhoo",
		TitleName: "达恩比杜",
		TitleCode: "b_dhanbidhoo",
	}
}
