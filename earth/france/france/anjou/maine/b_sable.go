package maine

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 萨布莱SableBarony struct {
	feud.BaseBarony
}

var BSable萨布莱 feud.Barony = &萨布莱SableBarony{}

func init() {
    f := BSable萨布莱.(*萨布莱SableBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "sable",
		TitleName: "萨布莱",
		TitleCode: "b_sable",
	}
}
