package shahrazur

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 卡卜甘KabganBarony struct {
	feud.BaseBarony
}

var BKabgan卡卜甘 feud.Barony = &卡卜甘KabganBarony{}

func init() {
	f := BKabgan卡卜甘.(*卡卜甘KabganBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kabgan",
		TitleName: "卡卜甘",
		TitleCode: "b_kabgan",
	}
}
