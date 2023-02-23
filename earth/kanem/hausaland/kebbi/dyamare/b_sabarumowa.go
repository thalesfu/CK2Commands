package dyamare

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 萨巴鲁莫瓦SabarumowaBarony struct {
	feud.BaseBarony
}

var BSabarumowa萨巴鲁莫瓦 feud.Barony = &萨巴鲁莫瓦SabarumowaBarony{}

func init() {
	f := BSabarumowa萨巴鲁莫瓦.(*萨巴鲁莫瓦SabarumowaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "sabarumowa",
		TitleName: "萨巴鲁莫瓦",
		TitleCode: "b_sabarumowa",
	}
}
