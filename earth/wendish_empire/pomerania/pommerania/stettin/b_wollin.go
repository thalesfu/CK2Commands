package stettin

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 沃林WollinBarony struct {
	feud.BaseBarony
}

var BWollin沃林 feud.Barony = &沃林WollinBarony{}

func init() {
    f := BWollin沃林.(*沃林WollinBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "wollin",
		TitleName: "沃林",
		TitleCode: "b_wollin",
	}
}
