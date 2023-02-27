package kyzyl

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 克孜勒KyzylBarony struct {
	feud.BaseBarony
}

var BKyzyl克孜勒 feud.Barony = &克孜勒KyzylBarony{}

func init() {
    f := BKyzyl克孜勒.(*克孜勒KyzylBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kyzyl",
		TitleName: "克孜勒",
		TitleCode: "b_kyzyl",
	}
}
