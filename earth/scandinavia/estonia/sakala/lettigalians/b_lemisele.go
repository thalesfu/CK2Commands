package lettigalians

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 莱米塞莱LemiseleBarony struct {
	feud.BaseBarony
}

var BLemisele莱米塞莱 feud.Barony = &莱米塞莱LemiseleBarony{}

func init() {
    f := BLemisele莱米塞莱.(*莱米塞莱LemiseleBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "lemisele",
		TitleName: "莱米塞莱",
		TitleCode: "b_lemisele",
	}
}
