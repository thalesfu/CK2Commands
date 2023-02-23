package sancerre

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 日安GienBarony struct {
	feud.BaseBarony
}

var BGien日安 feud.Barony = &日安GienBarony{}

func init() {
	f := BGien日安.(*日安GienBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "gien",
		TitleName: "日安",
		TitleCode: "b_gien",
	}
}
