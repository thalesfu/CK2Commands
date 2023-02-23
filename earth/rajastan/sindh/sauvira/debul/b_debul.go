package debul

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 提部罗DebulBarony struct {
	feud.BaseBarony
}

var BDebul提部罗 feud.Barony = &提部罗DebulBarony{}

func init() {
	f := BDebul提部罗.(*提部罗DebulBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "debul",
		TitleName: "提部罗",
		TitleCode: "b_debul",
	}
}
