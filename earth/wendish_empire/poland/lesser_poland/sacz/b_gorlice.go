package sacz

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 戈尔利采GorliceBarony struct {
	feud.BaseBarony
}

var BGorlice戈尔利采 feud.Barony = &戈尔利采GorliceBarony{}

func init() {
    f := BGorlice戈尔利采.(*戈尔利采GorliceBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "gorlice",
		TitleName: "戈尔利采",
		TitleCode: "b_gorlice",
	}
}
