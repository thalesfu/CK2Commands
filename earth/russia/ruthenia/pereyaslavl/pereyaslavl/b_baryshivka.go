package pereyaslavl

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 巴雷希夫卡BaryshivkaBarony struct {
	feud.BaseBarony
}

var BBaryshivka巴雷希夫卡 feud.Barony = &巴雷希夫卡BaryshivkaBarony{}

func init() {
	f := BBaryshivka巴雷希夫卡.(*巴雷希夫卡BaryshivkaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "baryshivka",
		TitleName: "巴雷希夫卡",
		TitleCode: "b_baryshivka",
	}
}
