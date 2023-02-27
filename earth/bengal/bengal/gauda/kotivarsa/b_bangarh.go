package kotivarsa

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 槃姞利呬BangarhBarony struct {
	feud.BaseBarony
}

var BBangarh槃姞利呬 feud.Barony = &槃姞利呬BangarhBarony{}

func init() {
    f := BBangarh槃姞利呬.(*槃姞利呬BangarhBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "bangarh",
		TitleName: "槃姞利呬",
		TitleCode: "b_bangarh",
	}
}
