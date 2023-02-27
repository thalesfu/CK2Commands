package montbeliard

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 埃蒂佩EtupesBarony struct {
	feud.BaseBarony
}

var BEtupes埃蒂佩 feud.Barony = &埃蒂佩EtupesBarony{}

func init() {
    f := BEtupes埃蒂佩.(*埃蒂佩EtupesBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "etupes",
		TitleName: "埃蒂佩",
		TitleCode: "b_etupes",
	}
}
