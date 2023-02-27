package asni

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 频陀吉BindkiBarony struct {
	feud.BaseBarony
}

var BBindki频陀吉 feud.Barony = &频陀吉BindkiBarony{}

func init() {
    f := BBindki频陀吉.(*频陀吉BindkiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "bindki",
		TitleName: "频陀吉",
		TitleCode: "b_bindki",
	}
}
