package trapezous

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿卢杰拉AlucraBarony struct {
	feud.BaseBarony
}

var BAlucra阿卢杰拉 feud.Barony = &阿卢杰拉AlucraBarony{}

func init() {
	f := BAlucra阿卢杰拉.(*阿卢杰拉AlucraBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "alucra",
		TitleName: "阿卢杰拉",
		TitleCode: "b_alucra",
	}
}
