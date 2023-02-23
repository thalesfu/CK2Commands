package janakpur

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 格莱亚KalaiyaBarony struct {
	feud.BaseBarony
}

var BKalaiya格莱亚 feud.Barony = &格莱亚KalaiyaBarony{}

func init() {
	f := BKalaiya格莱亚.(*格莱亚KalaiyaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kalaiya",
		TitleName: "格莱亚",
		TitleCode: "b_kalaiya",
	}
}
