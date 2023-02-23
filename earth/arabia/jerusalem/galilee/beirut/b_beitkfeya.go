package beirut

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 比克费亚BeitkfeyaBarony struct {
	feud.BaseBarony
}

var BBeitkfeya比克费亚 feud.Barony = &比克费亚BeitkfeyaBarony{}

func init() {
	f := BBeitkfeya比克费亚.(*比克费亚BeitkfeyaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "beitkfeya",
		TitleName: "比克费亚",
		TitleCode: "b_beitkfeya",
	}
}
