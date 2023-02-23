package veglia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 巴卡尔BakarBarony struct {
	feud.BaseBarony
}

var BBakar巴卡尔 feud.Barony = &巴卡尔BakarBarony{}

func init() {
	f := BBakar巴卡尔.(*巴卡尔BakarBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "bakar",
		TitleName: "巴卡尔",
		TitleCode: "b_bakar",
	}
}
