package hail

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿斯什纳AsshinanBarony struct {
	feud.BaseBarony
}

var BAsshinan阿斯什纳 feud.Barony = &阿斯什纳AsshinanBarony{}

func init() {
    f := BAsshinan阿斯什纳.(*阿斯什纳AsshinanBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "asshinan",
		TitleName: "阿斯什纳",
		TitleCode: "b_asshinan",
	}
}
