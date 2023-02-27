package medjybij

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 丘德尼夫ChudnivBarony struct {
	feud.BaseBarony
}

var BChudniv丘德尼夫 feud.Barony = &丘德尼夫ChudnivBarony{}

func init() {
    f := BChudniv丘德尼夫.(*丘德尼夫ChudnivBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "chudniv",
		TitleName: "丘德尼夫",
		TitleCode: "b_chudniv",
	}
}
