package usora

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 巴尼亚卢卡BanjalukaBarony struct {
	feud.BaseBarony
}

var BBanjaluka巴尼亚卢卡 feud.Barony = &巴尼亚卢卡BanjalukaBarony{}

func init() {
	f := BBanjaluka巴尼亚卢卡.(*巴尼亚卢卡BanjalukaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "banjaluka",
		TitleName: "巴尼亚卢卡",
		TitleCode: "b_banjaluka",
	}
}
