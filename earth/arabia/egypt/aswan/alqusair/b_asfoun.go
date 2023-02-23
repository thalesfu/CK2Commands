package alqusair

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿斯丰AsfounBarony struct {
	feud.BaseBarony
}

var BAsfoun阿斯丰 feud.Barony = &阿斯丰AsfounBarony{}

func init() {
	f := BAsfoun阿斯丰.(*阿斯丰AsfounBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "asfoun",
		TitleName: "阿斯丰",
		TitleCode: "b_asfoun",
	}
}
