package vozviahel

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 巴拉尼夫卡BaranivkaBarony struct {
	feud.BaseBarony
}

var BBaranivka巴拉尼夫卡 feud.Barony = &巴拉尼夫卡BaranivkaBarony{}

func init() {
    f := BBaranivka巴拉尼夫卡.(*巴拉尼夫卡BaranivkaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "baranivka",
		TitleName: "巴拉尼夫卡",
		TitleCode: "b_baranivka",
	}
}
