package kotivarsa

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 罗摩伐底RamavatiBarony struct {
	feud.BaseBarony
}

var BRamavati罗摩伐底 feud.Barony = &罗摩伐底RamavatiBarony{}

func init() {
    f := BRamavati罗摩伐底.(*罗摩伐底RamavatiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ramavati",
		TitleName: "罗摩伐底",
		TitleCode: "b_ramavati",
	}
}
