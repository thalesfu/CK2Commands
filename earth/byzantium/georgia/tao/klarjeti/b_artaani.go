package klarjeti

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿尔达汉ArtaaniBarony struct {
	feud.BaseBarony
}

var BArtaani阿尔达汉 feud.Barony = &阿尔达汉ArtaaniBarony{}

func init() {
    f := BArtaani阿尔达汉.(*阿尔达汉ArtaaniBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "artaani",
		TitleName: "阿尔达汉",
		TitleCode: "b_artaani",
	}
}
