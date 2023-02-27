package olomouc

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 城堡山HradisteBarony struct {
	feud.BaseBarony
}

var BHradiste城堡山 feud.Barony = &城堡山HradisteBarony{}

func init() {
    f := BHradiste城堡山.(*城堡山HradisteBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "hradiste",
		TitleName: "城堡山",
		TitleCode: "b_hradiste",
	}
}
