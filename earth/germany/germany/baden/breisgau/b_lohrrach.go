package breisgau

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 勒拉赫LohrrachBarony struct {
	feud.BaseBarony
}

var BLohrrach勒拉赫 feud.Barony = &勒拉赫LohrrachBarony{}

func init() {
	f := BLohrrach勒拉赫.(*勒拉赫LohrrachBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "lohrrach",
		TitleName: "勒拉赫",
		TitleCode: "b_lohrrach",
	}
}
