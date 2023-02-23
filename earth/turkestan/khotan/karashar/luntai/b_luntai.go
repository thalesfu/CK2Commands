package luntai

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 轮台LuntaiBarony struct {
	feud.BaseBarony
}

var BLuntai轮台 feud.Barony = &轮台LuntaiBarony{}

func init() {
	f := BLuntai轮台.(*轮台LuntaiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "luntai",
		TitleName: "轮台",
		TitleCode: "b_luntai",
	}
}
