package gurgi

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 商羯罗姞利呬ShankargarhBarony struct {
	feud.BaseBarony
}

var BShankargarh商羯罗姞利呬 feud.Barony = &商羯罗姞利呬ShankargarhBarony{}

func init() {
	f := BShankargarh商羯罗姞利呬.(*商羯罗姞利呬ShankargarhBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "shankargarh",
		TitleName: "商羯罗姞利呬",
		TitleCode: "b_shankargarh",
	}
}
