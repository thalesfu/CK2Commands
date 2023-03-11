package orekhovo_zouievo

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 穆尔米诺MurminoBarony struct {
	feud.BaseBarony
}

var BMurmino穆尔米诺 feud.Barony = &穆尔米诺MurminoBarony{}

func init() {
    f := BMurmino穆尔米诺.(*穆尔米诺MurminoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "murmino",
		TitleName: "穆尔米诺",
		TitleCode: "b_murmino",
	}
}
