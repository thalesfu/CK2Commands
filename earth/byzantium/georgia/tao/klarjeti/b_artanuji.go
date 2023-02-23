package klarjeti

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿尔达努奇ArtanujiBarony struct {
	feud.BaseBarony
}

var BArtanuji阿尔达努奇 feud.Barony = &阿尔达努奇ArtanujiBarony{}

func init() {
	f := BArtanuji阿尔达努奇.(*阿尔达努奇ArtanujiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "artanuji",
		TitleName: "阿尔达努奇",
		TitleCode: "b_artanuji",
	}
}
