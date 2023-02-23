package rottenburg

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 魏勒堡WeilerburgBarony struct {
	feud.BaseBarony
}

var BWeilerburg魏勒堡 feud.Barony = &魏勒堡WeilerburgBarony{}

func init() {
	f := BWeilerburg魏勒堡.(*魏勒堡WeilerburgBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "weilerburg",
		TitleName: "魏勒堡",
		TitleCode: "b_weilerburg",
	}
}
