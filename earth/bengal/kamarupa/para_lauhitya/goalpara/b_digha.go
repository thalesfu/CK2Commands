package goalpara

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 提格诃DighaBarony struct {
	feud.BaseBarony
}

var BDigha提格诃 feud.Barony = &提格诃DighaBarony{}

func init() {
    f := BDigha提格诃.(*提格诃DighaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "digha",
		TitleName: "提格诃",
		TitleCode: "b_digha",
	}
}
