package gar

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 昆莎KunsaBarony struct {
	feud.BaseBarony
}

var BKunsa昆莎 feud.Barony = &昆莎KunsaBarony{}

func init() {
	f := BKunsa昆莎.(*昆莎KunsaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kunsa",
		TitleName: "昆莎",
		TitleCode: "b_kunsa",
	}
}
