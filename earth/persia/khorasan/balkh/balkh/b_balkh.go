package balkh

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 缚喝BalkhBarony struct {
	feud.BaseBarony
}

var BBalkh缚喝 feud.Barony = &缚喝BalkhBarony{}

func init() {
	f := BBalkh缚喝.(*缚喝BalkhBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "balkh",
		TitleName: "缚喝",
		TitleCode: "b_balkh",
	}
}
