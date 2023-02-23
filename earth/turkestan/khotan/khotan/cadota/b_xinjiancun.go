package cadota

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 辛间村XinjiancunBarony struct {
	feud.BaseBarony
}

var BXinjiancun辛间村 feud.Barony = &辛间村XinjiancunBarony{}

func init() {
	f := BXinjiancun辛间村.(*辛间村XinjiancunBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "xinjiancun",
		TitleName: "辛间村",
		TitleCode: "b_xinjiancun",
	}
}
