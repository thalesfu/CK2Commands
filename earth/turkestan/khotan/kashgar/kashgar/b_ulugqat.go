package kashgar

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 乌恰UlugqatBarony struct {
	feud.BaseBarony
}

var BUlugqat乌恰 feud.Barony = &乌恰UlugqatBarony{}

func init() {
	f := BUlugqat乌恰.(*乌恰UlugqatBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ulugqat",
		TitleName: "乌恰",
		TitleCode: "b_ulugqat",
	}
}
