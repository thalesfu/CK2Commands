package boulogne

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 蒙特勒伊MontreuilBarony struct {
	feud.BaseBarony
}

var BMontreuil蒙特勒伊 feud.Barony = &蒙特勒伊MontreuilBarony{}

func init() {
	f := BMontreuil蒙特勒伊.(*蒙特勒伊MontreuilBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "montreuil",
		TitleName: "蒙特勒伊",
		TitleCode: "b_montreuil",
	}
}
