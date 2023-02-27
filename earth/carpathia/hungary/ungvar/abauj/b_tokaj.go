package abauj

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 托考伊TokajBarony struct {
	feud.BaseBarony
}

var BTokaj托考伊 feud.Barony = &托考伊TokajBarony{}

func init() {
    f := BTokaj托考伊.(*托考伊TokajBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "tokaj",
		TitleName: "托考伊",
		TitleCode: "b_tokaj",
	}
}
