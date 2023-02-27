package bost

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 格里什克GereshkBarony struct {
	feud.BaseBarony
}

var BGereshk格里什克 feud.Barony = &格里什克GereshkBarony{}

func init() {
    f := BGereshk格里什克.(*格里什克GereshkBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "gereshk",
		TitleName: "格里什克",
		TitleCode: "b_gereshk",
	}
}
