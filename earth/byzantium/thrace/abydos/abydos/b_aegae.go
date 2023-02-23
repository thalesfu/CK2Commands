package abydos

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 埃盖AegaeBarony struct {
	feud.BaseBarony
}

var BAegae埃盖 feud.Barony = &埃盖AegaeBarony{}

func init() {
	f := BAegae埃盖.(*埃盖AegaeBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "aegae",
		TitleName: "埃盖",
		TitleCode: "b_aegae",
	}
}
