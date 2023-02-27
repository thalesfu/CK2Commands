package goalpara

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿旄梨AmauliBarony struct {
	feud.BaseBarony
}

var BAmauli阿旄梨 feud.Barony = &阿旄梨AmauliBarony{}

func init() {
    f := BAmauli阿旄梨.(*阿旄梨AmauliBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "amauli",
		TitleName: "阿旄梨",
		TitleCode: "b_amauli",
	}
}
