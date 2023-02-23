package pest

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 小克勒什KiskorosBarony struct {
	feud.BaseBarony
}

var BKiskoros小克勒什 feud.Barony = &小克勒什KiskorosBarony{}

func init() {
	f := BKiskoros小克勒什.(*小克勒什KiskorosBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kiskoros",
		TitleName: "小克勒什",
		TitleCode: "b_kiskoros",
	}
}
