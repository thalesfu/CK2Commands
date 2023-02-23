package zoubtsov

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 巴塔科沃BatakovoBarony struct {
	feud.BaseBarony
}

var BBatakovo巴塔科沃 feud.Barony = &巴塔科沃BatakovoBarony{}

func init() {
	f := BBatakovo巴塔科沃.(*巴塔科沃BatakovoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "batakovo",
		TitleName: "巴塔科沃",
		TitleCode: "b_batakovo",
	}
}
