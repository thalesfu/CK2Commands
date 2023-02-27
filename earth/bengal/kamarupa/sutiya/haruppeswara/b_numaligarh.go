package haruppeswara

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 奴摩梨堡NumaligarhBarony struct {
	feud.BaseBarony
}

var BNumaligarh奴摩梨堡 feud.Barony = &奴摩梨堡NumaligarhBarony{}

func init() {
    f := BNumaligarh奴摩梨堡.(*奴摩梨堡NumaligarhBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "numaligarh",
		TitleName: "奴摩梨堡",
		TitleCode: "b_numaligarh",
	}
}
