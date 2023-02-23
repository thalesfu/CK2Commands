package tiberias

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 艾特拉AshteraBarony struct {
	feud.BaseBarony
}

var BAshtera艾特拉 feud.Barony = &艾特拉AshteraBarony{}

func init() {
	f := BAshtera艾特拉.(*艾特拉AshteraBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ashtera",
		TitleName: "艾特拉",
		TitleCode: "b_ashtera",
	}
}
