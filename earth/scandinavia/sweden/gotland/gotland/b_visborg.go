package gotland

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 维斯堡VisborgBarony struct {
	feud.BaseBarony
}

var BVisborg维斯堡 feud.Barony = &维斯堡VisborgBarony{}

func init() {
	f := BVisborg维斯堡.(*维斯堡VisborgBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "visborg",
		TitleName: "维斯堡",
		TitleCode: "b_visborg",
	}
}
