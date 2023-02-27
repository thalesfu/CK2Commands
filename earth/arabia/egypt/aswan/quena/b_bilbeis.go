package quena

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 比勒拜斯BilbeisBarony struct {
	feud.BaseBarony
}

var BBilbeis比勒拜斯 feud.Barony = &比勒拜斯BilbeisBarony{}

func init() {
    f := BBilbeis比勒拜斯.(*比勒拜斯BilbeisBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "bilbeis",
		TitleName: "比勒拜斯",
		TitleCode: "b_bilbeis",
	}
}
