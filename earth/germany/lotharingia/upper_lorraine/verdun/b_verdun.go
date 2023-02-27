package verdun

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 凡尔登VerdunBarony struct {
	feud.BaseBarony
}

var BVerdun凡尔登 feud.Barony = &凡尔登VerdunBarony{}

func init() {
    f := BVerdun凡尔登.(*凡尔登VerdunBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "verdun",
		TitleName: "凡尔登",
		TitleCode: "b_verdun",
	}
}
