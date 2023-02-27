package denia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿拉坎特AlicanteBarony struct {
	feud.BaseBarony
}

var BAlicante阿拉坎特 feud.Barony = &阿拉坎特AlicanteBarony{}

func init() {
    f := BAlicante阿拉坎特.(*阿拉坎特AlicanteBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "alicante",
		TitleName: "阿拉坎特",
		TitleCode: "b_alicante",
	}
}
