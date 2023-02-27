package soli

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 乌日采UsiceBarony struct {
	feud.BaseBarony
}

var BUsice乌日采 feud.Barony = &乌日采UsiceBarony{}

func init() {
    f := BUsice乌日采.(*乌日采UsiceBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "usice",
		TitleName: "乌日采",
		TitleCode: "b_usice",
	}
}
