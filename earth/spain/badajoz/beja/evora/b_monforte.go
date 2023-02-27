package evora

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 蒙福蒂MonforteBarony struct {
	feud.BaseBarony
}

var BMonforte蒙福蒂 feud.Barony = &蒙福蒂MonforteBarony{}

func init() {
    f := BMonforte蒙福蒂.(*蒙福蒂MonforteBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "monforte",
		TitleName: "蒙福蒂",
		TitleCode: "b_monforte",
	}
}
