package gowrie

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 珀斯PerthBarony struct {
	feud.BaseBarony
}

var BPerth珀斯 feud.Barony = &珀斯PerthBarony{}

func init() {
    f := BPerth珀斯.(*珀斯PerthBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "perth",
		TitleName: "珀斯",
		TitleCode: "b_perth",
	}
}
