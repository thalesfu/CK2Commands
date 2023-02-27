package buqtirma

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 朔克塔尔ShoktalBarony struct {
	feud.BaseBarony
}

var BShoktal朔克塔尔 feud.Barony = &朔克塔尔ShoktalBarony{}

func init() {
    f := BShoktal朔克塔尔.(*朔克塔尔ShoktalBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "shoktal",
		TitleName: "朔克塔尔",
		TitleCode: "b_shoktal",
	}
}
