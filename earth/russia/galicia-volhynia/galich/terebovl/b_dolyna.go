package terebovl

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 多利纳DolynaBarony struct {
	feud.BaseBarony
}

var BDolyna多利纳 feud.Barony = &多利纳DolynaBarony{}

func init() {
    f := BDolyna多利纳.(*多利纳DolynaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "dolyna",
		TitleName: "多利纳",
		TitleCode: "b_dolyna",
	}
}
