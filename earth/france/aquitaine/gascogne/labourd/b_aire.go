package labourd

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 艾尔AireBarony struct {
	feud.BaseBarony
}

var BAire艾尔 feud.Barony = &艾尔AireBarony{}

func init() {
    f := BAire艾尔.(*艾尔AireBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "aire",
		TitleName: "艾尔",
		TitleCode: "b_aire",
	}
}
