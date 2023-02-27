package al_bichri

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 锡尔SirhiBarony struct {
	feud.BaseBarony
}

var BSirhi锡尔 feud.Barony = &锡尔SirhiBarony{}

func init() {
    f := BSirhi锡尔.(*锡尔SirhiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "sirhi",
		TitleName: "锡尔",
		TitleCode: "b_sirhi",
	}
}
