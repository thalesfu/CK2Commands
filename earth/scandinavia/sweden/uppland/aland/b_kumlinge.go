package aland

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 库姆灵厄KumlingeBarony struct {
	feud.BaseBarony
}

var BKumlinge库姆灵厄 feud.Barony = &库姆灵厄KumlingeBarony{}

func init() {
	f := BKumlinge库姆灵厄.(*库姆灵厄KumlingeBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kumlinge",
		TitleName: "库姆灵厄",
		TitleCode: "b_kumlinge",
	}
}
