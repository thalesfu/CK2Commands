package bornholm

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 内克瑟NexoBarony struct {
	feud.BaseBarony
}

var BNexo内克瑟 feud.Barony = &内克瑟NexoBarony{}

func init() {
	f := BNexo内克瑟.(*内克瑟NexoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "nexo",
		TitleName: "内克瑟",
		TitleCode: "b_nexo",
	}
}
