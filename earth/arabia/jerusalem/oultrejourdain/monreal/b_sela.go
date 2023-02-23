package monreal

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 西拉SelaBarony struct {
	feud.BaseBarony
}

var BSela西拉 feud.Barony = &西拉SelaBarony{}

func init() {
	f := BSela西拉.(*西拉SelaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "sela",
		TitleName: "西拉",
		TitleCode: "b_sela",
	}
}
