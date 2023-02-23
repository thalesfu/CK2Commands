package tripuri

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 吠烂VelanBarony struct {
	feud.BaseBarony
}

var BVelan吠烂 feud.Barony = &吠烂VelanBarony{}

func init() {
	f := BVelan吠烂.(*吠烂VelanBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "velan",
		TitleName: "吠烂",
		TitleCode: "b_velan",
	}
}
