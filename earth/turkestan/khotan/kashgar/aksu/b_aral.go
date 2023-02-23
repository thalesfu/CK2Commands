package aksu

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿拉尔AralBarony struct {
	feud.BaseBarony
}

var BAral阿拉尔 feud.Barony = &阿拉尔AralBarony{}

func init() {
	f := BAral阿拉尔.(*阿拉尔AralBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "aral",
		TitleName: "阿拉尔",
		TitleCode: "b_aral",
	}
}
