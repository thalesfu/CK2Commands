package castelo_branco

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 萨布加尔SabugalBarony struct {
	feud.BaseBarony
}

var BSabugal萨布加尔 feud.Barony = &萨布加尔SabugalBarony{}

func init() {
    f := BSabugal萨布加尔.(*萨布加尔SabugalBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "sabugal",
		TitleName: "萨布加尔",
		TitleCode: "b_sabugal",
	}
}
