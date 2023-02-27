package krizevci

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 维特罗维蒂察ViroviticaBarony struct {
	feud.BaseBarony
}

var BVirovitica维特罗维蒂察 feud.Barony = &维特罗维蒂察ViroviticaBarony{}

func init() {
    f := BVirovitica维特罗维蒂察.(*维特罗维蒂察ViroviticaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "virovitica",
		TitleName: "维特罗维蒂察",
		TitleCode: "b_virovitica",
	}
}
