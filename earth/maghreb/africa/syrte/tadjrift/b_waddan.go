package tadjrift

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 韦丹WaddanBarony struct {
	feud.BaseBarony
}

var BWaddan韦丹 feud.Barony = &韦丹WaddanBarony{}

func init() {
	f := BWaddan韦丹.(*韦丹WaddanBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "waddan",
		TitleName: "韦丹",
		TitleCode: "b_waddan",
	}
}
