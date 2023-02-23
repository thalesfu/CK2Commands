package blekinge

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿瓦谢尔AvaskarBarony struct {
	feud.BaseBarony
}

var BAvaskar阿瓦谢尔 feud.Barony = &阿瓦谢尔AvaskarBarony{}

func init() {
	f := BAvaskar阿瓦谢尔.(*阿瓦谢尔AvaskarBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "avaskar",
		TitleName: "阿瓦谢尔",
		TitleCode: "b_avaskar",
	}
}
