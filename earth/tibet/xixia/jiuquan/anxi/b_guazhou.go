package anxi

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 瓜州GuazhouBarony struct {
	feud.BaseBarony
}

var BGuazhou瓜州 feud.Barony = &瓜州GuazhouBarony{}

func init() {
	f := BGuazhou瓜州.(*瓜州GuazhouBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "guazhou",
		TitleName: "瓜州",
		TitleCode: "b_guazhou",
	}
}
