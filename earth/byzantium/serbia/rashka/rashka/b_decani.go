package rashka

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 德查尼DecaniBarony struct {
	feud.BaseBarony
}

var BDecani德查尼 feud.Barony = &德查尼DecaniBarony{}

func init() {
	f := BDecani德查尼.(*德查尼DecaniBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "decani",
		TitleName: "德查尼",
		TitleCode: "b_decani",
	}
}
