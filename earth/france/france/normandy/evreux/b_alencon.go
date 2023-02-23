package evreux

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿朗松AlenconBarony struct {
	feud.BaseBarony
}

var BAlencon阿朗松 feud.Barony = &阿朗松AlenconBarony{}

func init() {
	f := BAlencon阿朗松.(*阿朗松AlenconBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "alencon",
		TitleName: "阿朗松",
		TitleCode: "b_alencon",
	}
}
