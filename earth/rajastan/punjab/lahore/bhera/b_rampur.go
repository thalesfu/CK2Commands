package bhera

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 罗摩补罗RampurBarony struct {
	feud.BaseBarony
}

var BRampur罗摩补罗 feud.Barony = &罗摩补罗RampurBarony{}

func init() {
	f := BRampur罗摩补罗.(*罗摩补罗RampurBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "rampur",
		TitleName: "罗摩补罗",
		TitleCode: "b_rampur",
	}
}
