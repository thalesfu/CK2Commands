package ludrava

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 罗那迦补罗RanakpurBarony struct {
	feud.BaseBarony
}

var BRanakpur罗那迦补罗 feud.Barony = &罗那迦补罗RanakpurBarony{}

func init() {
	f := BRanakpur罗那迦补罗.(*罗那迦补罗RanakpurBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ranakpur",
		TitleName: "罗那迦补罗",
		TitleCode: "b_ranakpur",
	}
}
