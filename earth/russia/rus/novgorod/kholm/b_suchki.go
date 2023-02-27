package kholm

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 苏奇基SuchkiBarony struct {
	feud.BaseBarony
}

var BSuchki苏奇基 feud.Barony = &苏奇基SuchkiBarony{}

func init() {
    f := BSuchki苏奇基.(*苏奇基SuchkiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "suchki",
		TitleName: "苏奇基",
		TitleCode: "b_suchki",
	}
}
