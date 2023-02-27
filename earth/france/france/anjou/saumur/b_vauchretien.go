package saumur

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 沃克雷蒂安VauchretienBarony struct {
	feud.BaseBarony
}

var BVauchretien沃克雷蒂安 feud.Barony = &沃克雷蒂安VauchretienBarony{}

func init() {
    f := BVauchretien沃克雷蒂安.(*沃克雷蒂安VauchretienBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "vauchretien",
		TitleName: "沃克雷蒂安",
		TitleCode: "b_vauchretien",
	}
}
