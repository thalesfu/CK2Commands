package chitrakut

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 罗摩补罗RampuraBarony struct {
	feud.BaseBarony
}

var BRampura罗摩补罗 feud.Barony = &罗摩补罗RampuraBarony{}

func init() {
	f := BRampura罗摩补罗.(*罗摩补罗RampuraBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "rampura",
		TitleName: "罗摩补罗",
		TitleCode: "b_rampura",
	}
}
