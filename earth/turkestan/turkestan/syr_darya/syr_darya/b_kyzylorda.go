package syr_darya

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 克孜勒奥尔达KyzylordaBarony struct {
	feud.BaseBarony
}

var BKyzylorda克孜勒奥尔达 feud.Barony = &克孜勒奥尔达KyzylordaBarony{}

func init() {
    f := BKyzylorda克孜勒奥尔达.(*克孜勒奥尔达KyzylordaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kyzylorda",
		TitleName: "克孜勒奥尔达",
		TitleCode: "b_kyzylorda",
	}
}
