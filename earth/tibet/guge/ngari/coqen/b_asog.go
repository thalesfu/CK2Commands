package coqen

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿索AsogBarony struct {
	feud.BaseBarony
}

var BAsog阿索 feud.Barony = &阿索AsogBarony{}

func init() {
    f := BAsog阿索.(*阿索AsogBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "asog",
		TitleName: "阿索",
		TitleCode: "b_asog",
	}
}
