package kiri

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 基塔KitaBarony struct {
	feud.BaseBarony
}

var BKita基塔 feud.Barony = &基塔KitaBarony{}

func init() {
	f := BKita基塔.(*基塔KitaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kita",
		TitleName: "基塔",
		TitleCode: "b_kita",
	}
}
