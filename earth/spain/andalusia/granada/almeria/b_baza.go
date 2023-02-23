package almeria

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 巴萨BazaBarony struct {
	feud.BaseBarony
}

var BBaza巴萨 feud.Barony = &巴萨BazaBarony{}

func init() {
	f := BBaza巴萨.(*巴萨BazaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "baza",
		TitleName: "巴萨",
		TitleCode: "b_baza",
	}
}
