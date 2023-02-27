package angermanland

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 索莱夫特奥SollefteaBarony struct {
	feud.BaseBarony
}

var BSolleftea索莱夫特奥 feud.Barony = &索莱夫特奥SollefteaBarony{}

func init() {
    f := BSolleftea索莱夫特奥.(*索莱夫特奥SollefteaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "solleftea",
		TitleName: "索莱夫特奥",
		TitleCode: "b_solleftea",
	}
}
