package parnakheta

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 半挐契吒ParnakhetaBarony struct {
	feud.BaseBarony
}

var BParnakheta半挐契吒 feud.Barony = &半挐契吒ParnakhetaBarony{}

func init() {
    f := BParnakheta半挐契吒.(*半挐契吒ParnakhetaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "parnakheta",
		TitleName: "半挐契吒",
		TitleCode: "b_parnakheta",
	}
}
