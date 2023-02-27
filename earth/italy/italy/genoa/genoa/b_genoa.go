package genoa

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 热那亚GenoaBarony struct {
	feud.BaseBarony
}

var BGenoa热那亚 feud.Barony = &热那亚GenoaBarony{}

func init() {
    f := BGenoa热那亚.(*热那亚GenoaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "genoa",
		TitleName: "热那亚",
		TitleCode: "b_genoa",
	}
}
