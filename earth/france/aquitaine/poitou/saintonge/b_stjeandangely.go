package saintonge

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 圣让当热利StjeandangelyBarony struct {
	feud.BaseBarony
}

var BStjeandangely圣让当热利 feud.Barony = &圣让当热利StjeandangelyBarony{}

func init() {
	f := BStjeandangely圣让当热利.(*圣让当热利StjeandangelyBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "stjeandangely",
		TitleName: "圣让当热利",
		TitleCode: "b_stjeandangely",
	}
}
