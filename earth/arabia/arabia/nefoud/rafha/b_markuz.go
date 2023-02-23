package rafha

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 马库兹MarkuzBarony struct {
	feud.BaseBarony
}

var BMarkuz马库兹 feud.Barony = &马库兹MarkuzBarony{}

func init() {
	f := BMarkuz马库兹.(*马库兹MarkuzBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "markuz",
		TitleName: "马库兹",
		TitleCode: "b_markuz",
	}
}
