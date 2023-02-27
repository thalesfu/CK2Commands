package furstenberg

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 希尔绍HirsauBarony struct {
	feud.BaseBarony
}

var BHirsau希尔绍 feud.Barony = &希尔绍HirsauBarony{}

func init() {
    f := BHirsau希尔绍.(*希尔绍HirsauBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "hirsau",
		TitleName: "希尔绍",
		TitleCode: "b_hirsau",
	}
}
