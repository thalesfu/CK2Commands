package godwad

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 婆卢蒂耶BarodiyaBarony struct {
	feud.BaseBarony
}

var BBarodiya婆卢蒂耶 feud.Barony = &婆卢蒂耶BarodiyaBarony{}

func init() {
	f := BBarodiya婆卢蒂耶.(*婆卢蒂耶BarodiyaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "barodiya",
		TitleName: "婆卢蒂耶",
		TitleCode: "b_barodiya",
	}
}
