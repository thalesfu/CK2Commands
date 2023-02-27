package riga

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 伊克什基莱IkskileBarony struct {
	feud.BaseBarony
}

var BIkskile伊克什基莱 feud.Barony = &伊克什基莱IkskileBarony{}

func init() {
    f := BIkskile伊克什基莱.(*伊克什基莱IkskileBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ikskile",
		TitleName: "伊克什基莱",
		TitleCode: "b_ikskile",
	}
}
