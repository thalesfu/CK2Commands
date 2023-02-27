package kota

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 罗摩姞利呬RamgarhBarony struct {
	feud.BaseBarony
}

var BRamgarh罗摩姞利呬 feud.Barony = &罗摩姞利呬RamgarhBarony{}

func init() {
    f := BRamgarh罗摩姞利呬.(*罗摩姞利呬RamgarhBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ramgarh",
		TitleName: "罗摩姞利呬",
		TitleCode: "b_ramgarh",
	}
}
