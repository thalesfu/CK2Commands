package skardu

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 摩度补罗Madhupur_skarduBarony struct {
	feud.BaseBarony
}

var BMadhupur_skardu摩度补罗 feud.Barony = &摩度补罗Madhupur_skarduBarony{}

func init() {
    f := BMadhupur_skardu摩度补罗.(*摩度补罗Madhupur_skarduBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "madhupur_skardu",
		TitleName: "摩度补罗",
		TitleCode: "b_madhupur_skardu",
	}
}
