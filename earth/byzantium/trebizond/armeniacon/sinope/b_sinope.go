package sinope

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 锡诺皮SinopeBarony struct {
	feud.BaseBarony
}

var BSinope锡诺皮 feud.Barony = &锡诺皮SinopeBarony{}

func init() {
	f := BSinope锡诺皮.(*锡诺皮SinopeBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "sinope",
		TitleName: "锡诺皮",
		TitleCode: "b_sinope",
	}
}
