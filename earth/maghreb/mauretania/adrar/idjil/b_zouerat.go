package idjil

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 祖韦拉特ZoueratBarony struct {
	feud.BaseBarony
}

var BZouerat祖韦拉特 feud.Barony = &祖韦拉特ZoueratBarony{}

func init() {
	f := BZouerat祖韦拉特.(*祖韦拉特ZoueratBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "zouerat",
		TitleName: "祖韦拉特",
		TitleCode: "b_zouerat",
	}
}
