package daylam

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 盖达尔GheydarBarony struct {
	feud.BaseBarony
}

var BGheydar盖达尔 feud.Barony = &盖达尔GheydarBarony{}

func init() {
    f := BGheydar盖达尔.(*盖达尔GheydarBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "gheydar",
		TitleName: "盖达尔",
		TitleCode: "b_gheydar",
	}
}
