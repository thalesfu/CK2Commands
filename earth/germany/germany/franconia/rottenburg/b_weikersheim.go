package rottenburg

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 魏克斯海姆WeikersheimBarony struct {
	feud.BaseBarony
}

var BWeikersheim魏克斯海姆 feud.Barony = &魏克斯海姆WeikersheimBarony{}

func init() {
	f := BWeikersheim魏克斯海姆.(*魏克斯海姆WeikersheimBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "weikersheim",
		TitleName: "魏克斯海姆",
		TitleCode: "b_weikersheim",
	}
}
