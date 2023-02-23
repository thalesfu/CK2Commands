package massat

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 齐达尼亚ZidaniaBarony struct {
	feud.BaseBarony
}

var BZidania齐达尼亚 feud.Barony = &齐达尼亚ZidaniaBarony{}

func init() {
	f := BZidania齐达尼亚.(*齐达尼亚ZidaniaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "zidania",
		TitleName: "齐达尼亚",
		TitleCode: "b_zidania",
	}
}
