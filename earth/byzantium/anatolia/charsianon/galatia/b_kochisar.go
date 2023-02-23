package galatia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 科奇希萨尔KochisarBarony struct {
	feud.BaseBarony
}

var BKochisar科奇希萨尔 feud.Barony = &科奇希萨尔KochisarBarony{}

func init() {
	f := BKochisar科奇希萨尔.(*科奇希萨尔KochisarBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kochisar",
		TitleName: "科奇希萨尔",
		TitleCode: "b_kochisar",
	}
}
