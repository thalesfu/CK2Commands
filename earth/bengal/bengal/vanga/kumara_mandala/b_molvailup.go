package kumara_mandala

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 牟毗卢MolvailupBarony struct {
	feud.BaseBarony
}

var BMolvailup牟毗卢 feud.Barony = &牟毗卢MolvailupBarony{}

func init() {
    f := BMolvailup牟毗卢.(*牟毗卢MolvailupBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "molvailup",
		TitleName: "牟毗卢",
		TitleCode: "b_molvailup",
	}
}
