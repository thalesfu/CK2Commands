package herjedalen

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 尤斯内达尔LjusnedalBarony struct {
	feud.BaseBarony
}

var BLjusnedal尤斯内达尔 feud.Barony = &尤斯内达尔LjusnedalBarony{}

func init() {
	f := BLjusnedal尤斯内达尔.(*尤斯内达尔LjusnedalBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ljusnedal",
		TitleName: "尤斯内达尔",
		TitleCode: "b_ljusnedal",
	}
}
