package halsingland

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 托斯塔TastaBarony struct {
	feud.BaseBarony
}

var BTasta托斯塔 feud.Barony = &托斯塔TastaBarony{}

func init() {
    f := BTasta托斯塔.(*托斯塔TastaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "tasta",
		TitleName: "托斯塔",
		TitleCode: "b_tasta",
	}
}
