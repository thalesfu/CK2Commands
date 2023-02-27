package srirangapatna

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 波那索戈PanasogeBarony struct {
	feud.BaseBarony
}

var BPanasoge波那索戈 feud.Barony = &波那索戈PanasogeBarony{}

func init() {
    f := BPanasoge波那索戈.(*波那索戈PanasogeBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "panasoge",
		TitleName: "波那索戈",
		TitleCode: "b_panasoge",
	}
}
