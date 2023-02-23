package manan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 楚库托利TchoukoutoliBarony struct {
	feud.BaseBarony
}

var BTchoukoutoli楚库托利 feud.Barony = &楚库托利TchoukoutoliBarony{}

func init() {
	f := BTchoukoutoli楚库托利.(*楚库托利TchoukoutoliBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "tchoukoutoli",
		TitleName: "楚库托利",
		TitleCode: "b_tchoukoutoli",
	}
}
