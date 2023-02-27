package yumen

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 报恩BaoenBarony struct {
	feud.BaseBarony
}

var BBaoen报恩 feud.Barony = &报恩BaoenBarony{}

func init() {
    f := BBaoen报恩.(*报恩BaoenBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "baoen",
		TitleName: "报恩",
		TitleCode: "b_baoen",
	}
}
