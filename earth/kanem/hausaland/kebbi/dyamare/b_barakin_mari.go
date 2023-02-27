package dyamare

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 巴拉金马里Barakin_mariBarony struct {
	feud.BaseBarony
}

var BBarakin_mari巴拉金马里 feud.Barony = &巴拉金马里Barakin_mariBarony{}

func init() {
    f := BBarakin_mari巴拉金马里.(*巴拉金马里Barakin_mariBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "barakin_mari",
		TitleName: "巴拉金马里",
		TitleCode: "b_barakin_mari",
	}
}
