package antalo

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 安茨基AntsokiaBarony struct {
	feud.BaseBarony
}

var BAntsokia安茨基 feud.Barony = &安茨基AntsokiaBarony{}

func init() {
    f := BAntsokia安茨基.(*安茨基AntsokiaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "antsokia",
		TitleName: "安茨基",
		TitleCode: "b_antsokia",
	}
}
