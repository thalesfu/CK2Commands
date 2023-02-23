package ili

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 科克塔尔KoktalBarony struct {
	feud.BaseBarony
}

var BKoktal科克塔尔 feud.Barony = &科克塔尔KoktalBarony{}

func init() {
	f := BKoktal科克塔尔.(*科克塔尔KoktalBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "koktal",
		TitleName: "科克塔尔",
		TitleCode: "b_koktal",
	}
}
