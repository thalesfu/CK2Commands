package krizevci

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 克里热夫齐KrizevciBarony struct {
	feud.BaseBarony
}

var BKrizevci克里热夫齐 feud.Barony = &克里热夫齐KrizevciBarony{}

func init() {
    f := BKrizevci克里热夫齐.(*克里热夫齐KrizevciBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "krizevci",
		TitleName: "克里热夫齐",
		TitleCode: "b_krizevci",
	}
}
