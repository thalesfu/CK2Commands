package mstislavl

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 德里宾DrybinBarony struct {
	feud.BaseBarony
}

var BDrybin德里宾 feud.Barony = &德里宾DrybinBarony{}

func init() {
	f := BDrybin德里宾.(*德里宾DrybinBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "drybin",
		TitleName: "德里宾",
		TitleCode: "b_drybin",
	}
}
