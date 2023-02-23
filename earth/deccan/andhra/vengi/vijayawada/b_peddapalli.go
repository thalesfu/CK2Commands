package vijayawada

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 佩达帕尔利PeddapalliBarony struct {
	feud.BaseBarony
}

var BPeddapalli佩达帕尔利 feud.Barony = &佩达帕尔利PeddapalliBarony{}

func init() {
	f := BPeddapalli佩达帕尔利.(*佩达帕尔利PeddapalliBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "peddapalli",
		TitleName: "佩达帕尔利",
		TitleCode: "b_peddapalli",
	}
}
