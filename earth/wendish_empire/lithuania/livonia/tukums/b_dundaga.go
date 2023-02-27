package tukums

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 敦达加DundagaBarony struct {
	feud.BaseBarony
}

var BDundaga敦达加 feud.Barony = &敦达加DundagaBarony{}

func init() {
    f := BDundaga敦达加.(*敦达加DundagaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "dundaga",
		TitleName: "敦达加",
		TitleCode: "b_dundaga",
	}
}
