package sambalpur

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 三婆罗补罗SambalpurBarony struct {
	feud.BaseBarony
}

var BSambalpur三婆罗补罗 feud.Barony = &三婆罗补罗SambalpurBarony{}

func init() {
    f := BSambalpur三婆罗补罗.(*三婆罗补罗SambalpurBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "sambalpur",
		TitleName: "三婆罗补罗",
		TitleCode: "b_sambalpur",
	}
}
