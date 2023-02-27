package innse_gall

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 基尔塔拉格伦KiltaraglenBarony struct {
	feud.BaseBarony
}

var BKiltaraglen基尔塔拉格伦 feud.Barony = &基尔塔拉格伦KiltaraglenBarony{}

func init() {
    f := BKiltaraglen基尔塔拉格伦.(*基尔塔拉格伦KiltaraglenBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kiltaraglen",
		TitleName: "基尔塔拉格伦",
		TitleCode: "b_kiltaraglen",
	}
}
