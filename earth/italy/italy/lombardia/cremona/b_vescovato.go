package cremona

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 韦斯科瓦托VescovatoBarony struct {
	feud.BaseBarony
}

var BVescovato韦斯科瓦托 feud.Barony = &韦斯科瓦托VescovatoBarony{}

func init() {
    f := BVescovato韦斯科瓦托.(*韦斯科瓦托VescovatoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "vescovato",
		TitleName: "韦斯科瓦托",
		TitleCode: "b_vescovato",
	}
}
