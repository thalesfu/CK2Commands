package amisos

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 法蒂森PhadisaneBarony struct {
	feud.BaseBarony
}

var BPhadisane法蒂森 feud.Barony = &法蒂森PhadisaneBarony{}

func init() {
	f := BPhadisane法蒂森.(*法蒂森PhadisaneBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "phadisane",
		TitleName: "法蒂森",
		TitleCode: "b_phadisane",
	}
}
