package tribandapura

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 陀卜诃罗DabthalaBarony struct {
	feud.BaseBarony
}

var BDabthala陀卜诃罗 feud.Barony = &陀卜诃罗DabthalaBarony{}

func init() {
    f := BDabthala陀卜诃罗.(*陀卜诃罗DabthalaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "dabthala",
		TitleName: "陀卜诃罗",
		TitleCode: "b_dabthala",
	}
}
