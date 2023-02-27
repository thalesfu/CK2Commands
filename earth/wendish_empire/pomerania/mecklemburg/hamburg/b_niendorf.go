package hamburg

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 宁多夫NiendorfBarony struct {
	feud.BaseBarony
}

var BNiendorf宁多夫 feud.Barony = &宁多夫NiendorfBarony{}

func init() {
    f := BNiendorf宁多夫.(*宁多夫NiendorfBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "niendorf",
		TitleName: "宁多夫",
		TitleCode: "b_niendorf",
	}
}
