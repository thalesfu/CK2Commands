package ratanpur

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 尸尔瓦尔ShirwalBarony struct {
	feud.BaseBarony
}

var BShirwal尸尔瓦尔 feud.Barony = &尸尔瓦尔ShirwalBarony{}

func init() {
	f := BShirwal尸尔瓦尔.(*尸尔瓦尔ShirwalBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "shirwal",
		TitleName: "尸尔瓦尔",
		TitleCode: "b_shirwal",
	}
}
