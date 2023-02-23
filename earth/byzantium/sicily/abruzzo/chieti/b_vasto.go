package chieti

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 瓦斯托VastoBarony struct {
	feud.BaseBarony
}

var BVasto瓦斯托 feud.Barony = &瓦斯托VastoBarony{}

func init() {
	f := BVasto瓦斯托.(*瓦斯托VastoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "vasto",
		TitleName: "瓦斯托",
		TitleCode: "b_vasto",
	}
}
