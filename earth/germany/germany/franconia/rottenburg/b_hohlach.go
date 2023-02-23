package rottenburg

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 霍拉赫HohlachBarony struct {
	feud.BaseBarony
}

var BHohlach霍拉赫 feud.Barony = &霍拉赫HohlachBarony{}

func init() {
	f := BHohlach霍拉赫.(*霍拉赫HohlachBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "hohlach",
		TitleName: "霍拉赫",
		TitleCode: "b_hohlach",
	}
}
