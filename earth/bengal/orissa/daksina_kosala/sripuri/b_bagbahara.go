package sripuri

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 巴格巴哈拉BagbaharaBarony struct {
	feud.BaseBarony
}

var BBagbahara巴格巴哈拉 feud.Barony = &巴格巴哈拉BagbaharaBarony{}

func init() {
    f := BBagbahara巴格巴哈拉.(*巴格巴哈拉BagbaharaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "bagbahara",
		TitleName: "巴格巴哈拉",
		TitleCode: "b_bagbahara",
	}
}
