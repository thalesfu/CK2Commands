package sieradzko-leczyckie

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 沃维奇LowiczBarony struct {
	feud.BaseBarony
}

var BLowicz沃维奇 feud.Barony = &沃维奇LowiczBarony{}

func init() {
    f := BLowicz沃维奇.(*沃维奇LowiczBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "lowicz",
		TitleName: "沃维奇",
		TitleCode: "b_lowicz",
	}
}
