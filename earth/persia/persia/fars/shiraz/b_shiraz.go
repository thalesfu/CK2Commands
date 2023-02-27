package shiraz

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 设拉子ShirazBarony struct {
	feud.BaseBarony
}

var BShiraz设拉子 feud.Barony = &设拉子ShirazBarony{}

func init() {
    f := BShiraz设拉子.(*设拉子ShirazBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "shiraz",
		TitleName: "设拉子",
		TitleCode: "b_shiraz",
	}
}
