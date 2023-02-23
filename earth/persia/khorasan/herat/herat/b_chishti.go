package herat

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 契斯提ChishtiBarony struct {
	feud.BaseBarony
}

var BChishti契斯提 feud.Barony = &契斯提ChishtiBarony{}

func init() {
	f := BChishti契斯提.(*契斯提ChishtiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "chishti",
		TitleName: "契斯提",
		TitleCode: "b_chishti",
	}
}
