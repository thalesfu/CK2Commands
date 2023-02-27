package vestfold

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 凯于庞KaupangBarony struct {
	feud.BaseBarony
}

var BKaupang凯于庞 feud.Barony = &凯于庞KaupangBarony{}

func init() {
    f := BKaupang凯于庞.(*凯于庞KaupangBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kaupang",
		TitleName: "凯于庞",
		TitleCode: "b_kaupang",
	}
}
