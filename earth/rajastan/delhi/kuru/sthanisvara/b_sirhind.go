package sthanisvara

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 西尔欣德SirhindBarony struct {
	feud.BaseBarony
}

var BSirhind西尔欣德 feud.Barony = &西尔欣德SirhindBarony{}

func init() {
    f := BSirhind西尔欣德.(*西尔欣德SirhindBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "sirhind",
		TitleName: "西尔欣德",
		TitleCode: "b_sirhind",
	}
}
