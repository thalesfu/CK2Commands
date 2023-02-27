package palermo

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 米斯特雷塔MistrettaBarony struct {
	feud.BaseBarony
}

var BMistretta米斯特雷塔 feud.Barony = &米斯特雷塔MistrettaBarony{}

func init() {
    f := BMistretta米斯特雷塔.(*米斯特雷塔MistrettaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "mistretta",
		TitleName: "米斯特雷塔",
		TitleCode: "b_mistretta",
	}
}
