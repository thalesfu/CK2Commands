package lhunzhub

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿朗NgarnangBarony struct {
	feud.BaseBarony
}

var BNgarnang阿朗 feud.Barony = &阿朗NgarnangBarony{}

func init() {
	f := BNgarnang阿朗.(*阿朗NgarnangBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ngarnang",
		TitleName: "阿朗",
		TitleCode: "b_ngarnang",
	}
}
