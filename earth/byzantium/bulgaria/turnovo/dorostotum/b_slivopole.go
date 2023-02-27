package dorostotum

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 斯利沃波莱SlivopoleBarony struct {
	feud.BaseBarony
}

var BSlivopole斯利沃波莱 feud.Barony = &斯利沃波莱SlivopoleBarony{}

func init() {
    f := BSlivopole斯利沃波莱.(*斯利沃波莱SlivopoleBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "slivopole",
		TitleName: "斯利沃波莱",
		TitleCode: "b_slivopole",
	}
}
