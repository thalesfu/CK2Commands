package samatata

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 提婆钵伐多DevaparvataBarony struct {
	feud.BaseBarony
}

var BDevaparvata提婆钵伐多 feud.Barony = &提婆钵伐多DevaparvataBarony{}

func init() {
	f := BDevaparvata提婆钵伐多.(*提婆钵伐多DevaparvataBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "devaparvata",
		TitleName: "提婆钵伐多",
		TitleCode: "b_devaparvata",
	}
}
