package sravasti

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 婆卢醯耶BarohiyaBarony struct {
	feud.BaseBarony
}

var BBarohiya婆卢醯耶 feud.Barony = &婆卢醯耶BarohiyaBarony{}

func init() {
	f := BBarohiya婆卢醯耶.(*婆卢醯耶BarohiyaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "barohiya",
		TitleName: "婆卢醯耶",
		TitleCode: "b_barohiya",
	}
}
