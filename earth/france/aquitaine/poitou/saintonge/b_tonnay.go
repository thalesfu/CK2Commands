package saintonge

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 托奈TonnayBarony struct {
	feud.BaseBarony
}

var BTonnay托奈 feud.Barony = &托奈TonnayBarony{}

func init() {
	f := BTonnay托奈.(*托奈TonnayBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "tonnay",
		TitleName: "托奈",
		TitleCode: "b_tonnay",
	}
}
