package yopurga

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 八里庄PalichuangBarony struct {
	feud.BaseBarony
}

var BPalichuang八里庄 feud.Barony = &八里庄PalichuangBarony{}

func init() {
    f := BPalichuang八里庄.(*八里庄PalichuangBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "palichuang",
		TitleName: "八里庄",
		TitleCode: "b_palichuang",
	}
}
