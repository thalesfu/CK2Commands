package leicester

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 纽斯特德NewsteadBarony struct {
	feud.BaseBarony
}

var BNewstead纽斯特德 feud.Barony = &纽斯特德NewsteadBarony{}

func init() {
	f := BNewstead纽斯特德.(*纽斯特德NewsteadBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "newstead",
		TitleName: "纽斯特德",
		TitleCode: "b_newstead",
	}
}
