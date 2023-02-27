package dublin

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 克朗多金ClondalkinBarony struct {
	feud.BaseBarony
}

var BClondalkin克朗多金 feud.Barony = &克朗多金ClondalkinBarony{}

func init() {
    f := BClondalkin克朗多金.(*克朗多金ClondalkinBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "clondalkin",
		TitleName: "克朗多金",
		TitleCode: "b_clondalkin",
	}
}
