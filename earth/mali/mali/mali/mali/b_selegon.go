package mali

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 塞莱贡SelegonBarony struct {
	feud.BaseBarony
}

var BSelegon塞莱贡 feud.Barony = &塞莱贡SelegonBarony{}

func init() {
	f := BSelegon塞莱贡.(*塞莱贡SelegonBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "selegon",
		TitleName: "塞莱贡",
		TitleCode: "b_selegon",
	}
}
