package auvergne

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 图尔纳尔TournoelBarony struct {
	feud.BaseBarony
}

var BTournoel图尔纳尔 feud.Barony = &图尔纳尔TournoelBarony{}

func init() {
	f := BTournoel图尔纳尔.(*图尔纳尔TournoelBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "tournoel",
		TitleName: "图尔纳尔",
		TitleCode: "b_tournoel",
	}
}
