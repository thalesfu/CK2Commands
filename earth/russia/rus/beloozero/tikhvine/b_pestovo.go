package tikhvine

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 佩斯托沃PestovoBarony struct {
	feud.BaseBarony
}

var BPestovo佩斯托沃 feud.Barony = &佩斯托沃PestovoBarony{}

func init() {
    f := BPestovo佩斯托沃.(*佩斯托沃PestovoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "pestovo",
		TitleName: "佩斯托沃",
		TitleCode: "b_pestovo",
	}
}
