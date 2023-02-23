package rosello

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 卡尼戈CanigoBarony struct {
	feud.BaseBarony
}

var BCanigo卡尼戈 feud.Barony = &卡尼戈CanigoBarony{}

func init() {
	f := BCanigo卡尼戈.(*卡尼戈CanigoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "canigo",
		TitleName: "卡尼戈",
		TitleCode: "b_canigo",
	}
}
