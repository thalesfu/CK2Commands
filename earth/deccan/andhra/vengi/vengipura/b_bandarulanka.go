package vengipura

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 班达鲁兰卡BandarulankaBarony struct {
	feud.BaseBarony
}

var BBandarulanka班达鲁兰卡 feud.Barony = &班达鲁兰卡BandarulankaBarony{}

func init() {
    f := BBandarulanka班达鲁兰卡.(*班达鲁兰卡BandarulankaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "bandarulanka",
		TitleName: "班达鲁兰卡",
		TitleCode: "b_bandarulanka",
	}
}
