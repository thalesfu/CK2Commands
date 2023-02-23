package herat

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 亦鲁HeratBarony struct {
	feud.BaseBarony
}

var BHerat亦鲁 feud.Barony = &亦鲁HeratBarony{}

func init() {
	f := BHerat亦鲁.(*亦鲁HeratBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "herat",
		TitleName: "亦鲁",
		TitleCode: "b_herat",
	}
}
