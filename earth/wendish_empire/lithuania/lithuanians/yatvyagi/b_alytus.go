package yatvyagi

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿利图斯AlytusBarony struct {
	feud.BaseBarony
}

var BAlytus阿利图斯 feud.Barony = &阿利图斯AlytusBarony{}

func init() {
    f := BAlytus阿利图斯.(*阿利图斯AlytusBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "alytus",
		TitleName: "阿利图斯",
		TitleCode: "b_alytus",
	}
}
