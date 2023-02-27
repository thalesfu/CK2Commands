package gabes

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 哈德贾HaddejBarony struct {
	feud.BaseBarony
}

var BHaddej哈德贾 feud.Barony = &哈德贾HaddejBarony{}

func init() {
    f := BHaddej哈德贾.(*哈德贾HaddejBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "haddej",
		TitleName: "哈德贾",
		TitleCode: "b_haddej",
	}
}
