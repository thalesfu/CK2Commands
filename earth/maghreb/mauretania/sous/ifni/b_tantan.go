package ifni

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 坦坦TantanBarony struct {
	feud.BaseBarony
}

var BTantan坦坦 feud.Barony = &坦坦TantanBarony{}

func init() {
    f := BTantan坦坦.(*坦坦TantanBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "tantan",
		TitleName: "坦坦",
		TitleCode: "b_tantan",
	}
}
