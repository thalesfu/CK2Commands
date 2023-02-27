package pisa

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 圣米尼亚托SanminiatoBarony struct {
	feud.BaseBarony
}

var BSanminiato圣米尼亚托 feud.Barony = &圣米尼亚托SanminiatoBarony{}

func init() {
    f := BSanminiato圣米尼亚托.(*圣米尼亚托SanminiatoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "sanminiato",
		TitleName: "圣米尼亚托",
		TitleCode: "b_sanminiato",
	}
}
