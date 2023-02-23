package kebbi

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿尔贡古ArgunguBarony struct {
	feud.BaseBarony
}

var BArgungu阿尔贡古 feud.Barony = &阿尔贡古ArgunguBarony{}

func init() {
	f := BArgungu阿尔贡古.(*阿尔贡古ArgunguBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "argungu",
		TitleName: "阿尔贡古",
		TitleCode: "b_argungu",
	}
}
