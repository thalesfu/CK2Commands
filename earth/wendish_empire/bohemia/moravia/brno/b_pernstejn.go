package brno

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 佩恩什泰因PernstejnBarony struct {
	feud.BaseBarony
}

var BPernstejn佩恩什泰因 feud.Barony = &佩恩什泰因PernstejnBarony{}

func init() {
    f := BPernstejn佩恩什泰因.(*佩恩什泰因PernstejnBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "pernstejn",
		TitleName: "佩恩什泰因",
		TitleCode: "b_pernstejn",
	}
}
