package sjaelland

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 雷勒LejreBarony struct {
	feud.BaseBarony
}

var BLejre雷勒 feud.Barony = &雷勒LejreBarony{}

func init() {
    f := BLejre雷勒.(*雷勒LejreBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "lejre",
		TitleName: "雷勒",
		TitleCode: "b_lejre",
	}
}
