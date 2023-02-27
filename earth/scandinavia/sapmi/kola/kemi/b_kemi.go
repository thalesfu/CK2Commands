package kemi

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 凯米KemiBarony struct {
	feud.BaseBarony
}

var BKemi凯米 feud.Barony = &凯米KemiBarony{}

func init() {
    f := BKemi凯米.(*凯米KemiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kemi",
		TitleName: "凯米",
		TitleCode: "b_kemi",
	}
}
