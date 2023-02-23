package chagai

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 达德DadBarony struct {
	feud.BaseBarony
}

var BDad达德 feud.Barony = &达德DadBarony{}

func init() {
	f := BDad达德.(*达德DadBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "dad",
		TitleName: "达德",
		TitleCode: "b_dad",
	}
}
