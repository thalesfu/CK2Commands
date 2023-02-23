package sticht

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 武尔登WoerdenBarony struct {
	feud.BaseBarony
}

var BWoerden武尔登 feud.Barony = &武尔登WoerdenBarony{}

func init() {
	f := BWoerden武尔登.(*武尔登WoerdenBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "woerden",
		TitleName: "武尔登",
		TitleCode: "b_woerden",
	}
}
