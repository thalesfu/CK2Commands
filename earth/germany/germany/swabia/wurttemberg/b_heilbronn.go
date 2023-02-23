package wurttemberg

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 海尔布隆HeilbronnBarony struct {
	feud.BaseBarony
}

var BHeilbronn海尔布隆 feud.Barony = &海尔布隆HeilbronnBarony{}

func init() {
	f := BHeilbronn海尔布隆.(*海尔布隆HeilbronnBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "heilbronn",
		TitleName: "海尔布隆",
		TitleCode: "b_heilbronn",
	}
}
