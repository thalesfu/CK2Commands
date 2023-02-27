package nandagiri

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 军荼尼KundaniBarony struct {
	feud.BaseBarony
}

var BKundani军荼尼 feud.Barony = &军荼尼KundaniBarony{}

func init() {
    f := BKundani军荼尼.(*军荼尼KundaniBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kundani",
		TitleName: "军荼尼",
		TitleCode: "b_kundani",
	}
}
