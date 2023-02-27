package vanema

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 瓦内马VanemaBarony struct {
	feud.BaseBarony
}

var BVanema瓦内马 feud.Barony = &瓦内马VanemaBarony{}

func init() {
    f := BVanema瓦内马.(*瓦内马VanemaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "vanema",
		TitleName: "瓦内马",
		TitleCode: "b_vanema",
	}
}
