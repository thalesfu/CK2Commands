package abosyssel

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 斯坎讷堡SkanderborgBarony struct {
	feud.BaseBarony
}

var BSkanderborg斯坎讷堡 feud.Barony = &斯坎讷堡SkanderborgBarony{}

func init() {
	f := BSkanderborg斯坎讷堡.(*斯坎讷堡SkanderborgBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "skanderborg",
		TitleName: "斯坎讷堡",
		TitleCode: "b_skanderborg",
	}
}
