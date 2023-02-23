package ashir

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 西迪赫格SidihegBarony struct {
	feud.BaseBarony
}

var BSidiheg西迪赫格 feud.Barony = &西迪赫格SidihegBarony{}

func init() {
	f := BSidiheg西迪赫格.(*西迪赫格SidihegBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "sidiheg",
		TitleName: "西迪赫格",
		TitleCode: "b_sidiheg",
	}
}
