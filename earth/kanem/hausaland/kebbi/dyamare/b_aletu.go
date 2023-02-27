package dyamare

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿勒图AletuBarony struct {
	feud.BaseBarony
}

var BAletu阿勒图 feud.Barony = &阿勒图AletuBarony{}

func init() {
    f := BAletu阿勒图.(*阿勒图AletuBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "aletu",
		TitleName: "阿勒图",
		TitleCode: "b_aletu",
	}
}
