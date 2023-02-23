package albania

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 叶木里EmliBarony struct {
	feud.BaseBarony
}

var BEmli叶木里 feud.Barony = &叶木里EmliBarony{}

func init() {
	f := BEmli叶木里.(*叶木里EmliBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "emli",
		TitleName: "叶木里",
		TitleCode: "b_emli",
	}
}
