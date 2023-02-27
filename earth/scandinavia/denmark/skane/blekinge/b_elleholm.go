package blekinge

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 埃勒霍尔姆ElleholmBarony struct {
	feud.BaseBarony
}

var BElleholm埃勒霍尔姆 feud.Barony = &埃勒霍尔姆ElleholmBarony{}

func init() {
    f := BElleholm埃勒霍尔姆.(*埃勒霍尔姆ElleholmBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "elleholm",
		TitleName: "埃勒霍尔姆",
		TitleCode: "b_elleholm",
	}
}
