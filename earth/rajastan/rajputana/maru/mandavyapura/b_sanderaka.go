package mandavyapura

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 桑啼罗迦SanderakaBarony struct {
	feud.BaseBarony
}

var BSanderaka桑啼罗迦 feud.Barony = &桑啼罗迦SanderakaBarony{}

func init() {
	f := BSanderaka桑啼罗迦.(*桑啼罗迦SanderakaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "sanderaka",
		TitleName: "桑啼罗迦",
		TitleCode: "b_sanderaka",
	}
}
