package ajayameru

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 檀努DhanopBarony struct {
	feud.BaseBarony
}

var BDhanop檀努 feud.Barony = &檀努DhanopBarony{}

func init() {
    f := BDhanop檀努.(*檀努DhanopBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "dhanop",
		TitleName: "檀努",
		TitleCode: "b_dhanop",
	}
}
