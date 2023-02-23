package kerak

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 普嫩PunonBarony struct {
	feud.BaseBarony
}

var BPunon普嫩 feud.Barony = &普嫩PunonBarony{}

func init() {
	f := BPunon普嫩.(*普嫩PunonBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "punon",
		TitleName: "普嫩",
		TitleCode: "b_punon",
	}
}
