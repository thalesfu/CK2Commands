package konjikala

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 孔吉卡拉KonjikalaBarony struct {
	feud.BaseBarony
}

var BKonjikala孔吉卡拉 feud.Barony = &孔吉卡拉KonjikalaBarony{}

func init() {
    f := BKonjikala孔吉卡拉.(*孔吉卡拉KonjikalaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "konjikala",
		TitleName: "孔吉卡拉",
		TitleCode: "b_konjikala",
	}
}
