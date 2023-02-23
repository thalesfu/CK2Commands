package buzachi

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 杜尔涅瓦DurnevaBarony struct {
	feud.BaseBarony
}

var BDurneva杜尔涅瓦 feud.Barony = &杜尔涅瓦DurnevaBarony{}

func init() {
	f := BDurneva杜尔涅瓦.(*杜尔涅瓦DurnevaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "durneva",
		TitleName: "杜尔涅瓦",
		TitleCode: "b_durneva",
	}
}
