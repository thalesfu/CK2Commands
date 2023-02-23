package bayda

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 提姆纳TimnaBarony struct {
	feud.BaseBarony
}

var BTimna提姆纳 feud.Barony = &提姆纳TimnaBarony{}

func init() {
	f := BTimna提姆纳.(*提姆纳TimnaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "timna",
		TitleName: "提姆纳",
		TitleCode: "b_timna",
	}
}
