package trapezous

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 佩佩尔塔PaipertaBarony struct {
	feud.BaseBarony
}

var BPaiperta佩佩尔塔 feud.Barony = &佩佩尔塔PaipertaBarony{}

func init() {
	f := BPaiperta佩佩尔塔.(*佩佩尔塔PaipertaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "paiperta",
		TitleName: "佩佩尔塔",
		TitleCode: "b_paiperta",
	}
}
