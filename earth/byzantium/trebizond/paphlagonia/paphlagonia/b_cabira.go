package paphlagonia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 加比拉CabiraBarony struct {
	feud.BaseBarony
}

var BCabira加比拉 feud.Barony = &加比拉CabiraBarony{}

func init() {
	f := BCabira加比拉.(*加比拉CabiraBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "cabira",
		TitleName: "加比拉",
		TitleCode: "b_cabira",
	}
}
