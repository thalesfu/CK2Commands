package labourd

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿斯帕伦HasparrenBarony struct {
	feud.BaseBarony
}

var BHasparren阿斯帕伦 feud.Barony = &阿斯帕伦HasparrenBarony{}

func init() {
	f := BHasparren阿斯帕伦.(*阿斯帕伦HasparrenBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "hasparren",
		TitleName: "阿斯帕伦",
		TitleCode: "b_hasparren",
	}
}
