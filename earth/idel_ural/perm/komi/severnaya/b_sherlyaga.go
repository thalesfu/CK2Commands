package severnaya

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 舍尔利亚加SherlyagaBarony struct {
	feud.BaseBarony
}

var BSherlyaga舍尔利亚加 feud.Barony = &舍尔利亚加SherlyagaBarony{}

func init() {
    f := BSherlyaga舍尔利亚加.(*舍尔利亚加SherlyagaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "sherlyaga",
		TitleName: "舍尔利亚加",
		TitleCode: "b_sherlyaga",
	}
}
