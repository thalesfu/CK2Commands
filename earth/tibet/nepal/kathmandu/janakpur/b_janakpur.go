package janakpur

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阇那迦补罗JanakpurBarony struct {
	feud.BaseBarony
}

var BJanakpur阇那迦补罗 feud.Barony = &阇那迦补罗JanakpurBarony{}

func init() {
	f := BJanakpur阇那迦补罗.(*阇那迦补罗JanakpurBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "janakpur",
		TitleName: "阇那迦补罗",
		TitleCode: "b_janakpur",
	}
}
