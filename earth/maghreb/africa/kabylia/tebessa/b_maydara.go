package tebessa

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 迈德雷MaydaraBarony struct {
	feud.BaseBarony
}

var BMaydara迈德雷 feud.Barony = &迈德雷MaydaraBarony{}

func init() {
	f := BMaydara迈德雷.(*迈德雷MaydaraBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "maydara",
		TitleName: "迈德雷",
		TitleCode: "b_maydara",
	}
}
