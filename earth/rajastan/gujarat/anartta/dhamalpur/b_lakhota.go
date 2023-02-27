package dhamalpur

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 罗诃陀LakhotaBarony struct {
	feud.BaseBarony
}

var BLakhota罗诃陀 feud.Barony = &罗诃陀LakhotaBarony{}

func init() {
    f := BLakhota罗诃陀.(*罗诃陀LakhotaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "lakhota",
		TitleName: "罗诃陀",
		TitleCode: "b_lakhota",
	}
}
