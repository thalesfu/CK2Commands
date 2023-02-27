package jersika

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿索特AsoteBarony struct {
	feud.BaseBarony
}

var BAsote阿索特 feud.Barony = &阿索特AsoteBarony{}

func init() {
    f := BAsote阿索特.(*阿索特AsoteBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "asote",
		TitleName: "阿索特",
		TitleCode: "b_asote",
	}
}
