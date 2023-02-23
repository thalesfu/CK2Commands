package arta

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿尔塔ArtaBarony struct {
	feud.BaseBarony
}

var BArta阿尔塔 feud.Barony = &阿尔塔ArtaBarony{}

func init() {
	f := BArta阿尔塔.(*阿尔塔ArtaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "arta",
		TitleName: "阿尔塔",
		TitleCode: "b_arta",
	}
}
