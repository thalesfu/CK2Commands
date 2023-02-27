package uda

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 结雅ZeyaBarony struct {
	feud.BaseBarony
}

var BZeya结雅 feud.Barony = &结雅ZeyaBarony{}

func init() {
    f := BZeya结雅.(*结雅ZeyaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "zeya",
		TitleName: "结雅",
		TitleCode: "b_zeya",
	}
}
