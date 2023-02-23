package beirut

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 迈什盖拉MashgarahBarony struct {
	feud.BaseBarony
}

var BMashgarah迈什盖拉 feud.Barony = &迈什盖拉MashgarahBarony{}

func init() {
	f := BMashgarah迈什盖拉.(*迈什盖拉MashgarahBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "mashgarah",
		TitleName: "迈什盖拉",
		TitleCode: "b_mashgarah",
	}
}
