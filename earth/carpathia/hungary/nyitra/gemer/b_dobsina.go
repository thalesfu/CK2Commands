package gemer

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 多布希瑙DobsinaBarony struct {
	feud.BaseBarony
}

var BDobsina多布希瑙 feud.Barony = &多布希瑙DobsinaBarony{}

func init() {
    f := BDobsina多布希瑙.(*多布希瑙DobsinaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "dobsina",
		TitleName: "多布希瑙",
		TitleCode: "b_dobsina",
	}
}
