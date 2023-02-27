package vashka

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 列贾RezyaBarony struct {
	feud.BaseBarony
}

var BRezya列贾 feud.Barony = &列贾RezyaBarony{}

func init() {
    f := BRezya列贾.(*列贾RezyaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "rezya",
		TitleName: "列贾",
		TitleCode: "b_rezya",
	}
}
