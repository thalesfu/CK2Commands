package tyrone

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 科尔雷恩ColeraineBarony struct {
	feud.BaseBarony
}

var BColeraine科尔雷恩 feud.Barony = &科尔雷恩ColeraineBarony{}

func init() {
	f := BColeraine科尔雷恩.(*科尔雷恩ColeraineBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "coleraine",
		TitleName: "科尔雷恩",
		TitleCode: "b_coleraine",
	}
}
