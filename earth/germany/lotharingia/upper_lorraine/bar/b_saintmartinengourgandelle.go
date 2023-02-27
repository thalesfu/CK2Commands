package bar

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 古尔冈德勒的圣马尔坦SaintmartinengourgandelleBarony struct {
	feud.BaseBarony
}

var BSaintmartinengourgandelle古尔冈德勒的圣马尔坦 feud.Barony = &古尔冈德勒的圣马尔坦SaintmartinengourgandelleBarony{}

func init() {
    f := BSaintmartinengourgandelle古尔冈德勒的圣马尔坦.(*古尔冈德勒的圣马尔坦SaintmartinengourgandelleBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "saintmartinengourgandelle",
		TitleName: "古尔冈德勒的圣马尔坦",
		TitleCode: "b_saintmartinengourgandelle",
	}
}
