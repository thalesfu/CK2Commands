package chanderi

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 奢罗迦摩斯Jarai_ka_mathBarony struct {
	feud.BaseBarony
}

var BJarai_ka_math奢罗迦摩斯 feud.Barony = &奢罗迦摩斯Jarai_ka_mathBarony{}

func init() {
    f := BJarai_ka_math奢罗迦摩斯.(*奢罗迦摩斯Jarai_ka_mathBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "jarai_ka_math",
		TitleName: "奢罗迦摩斯",
		TitleCode: "b_jarai_ka_math",
	}
}
