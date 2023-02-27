package sthanisvara

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 舍甘婆利提毗居处Shakumbhri_deviBarony struct {
	feud.BaseBarony
}

var BShakumbhri_devi舍甘婆利提毗居处 feud.Barony = &舍甘婆利提毗居处Shakumbhri_deviBarony{}

func init() {
    f := BShakumbhri_devi舍甘婆利提毗居处.(*舍甘婆利提毗居处Shakumbhri_deviBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "shakumbhri_devi",
		TitleName: "舍甘婆利提毗居处",
		TitleCode: "b_shakumbhri_devi",
	}
}
