package skardu

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 希尔迈特HilmatBarony struct {
	feud.BaseBarony
}

var BHilmat希尔迈特 feud.Barony = &希尔迈特HilmatBarony{}

func init() {
    f := BHilmat希尔迈特.(*希尔迈特HilmatBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "hilmat",
		TitleName: "希尔迈特",
		TitleCode: "b_hilmat",
	}
}
