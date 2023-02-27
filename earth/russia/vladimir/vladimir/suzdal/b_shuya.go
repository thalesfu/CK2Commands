package suzdal

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 舒亚ShuyaBarony struct {
	feud.BaseBarony
}

var BShuya舒亚 feud.Barony = &舒亚ShuyaBarony{}

func init() {
    f := BShuya舒亚.(*舒亚ShuyaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "shuya",
		TitleName: "舒亚",
		TitleCode: "b_shuya",
	}
}
