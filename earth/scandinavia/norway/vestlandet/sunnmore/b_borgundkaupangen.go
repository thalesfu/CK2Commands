package sunnmore

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 博尔贡凯于庞恩BorgundkaupangenBarony struct {
	feud.BaseBarony
}

var BBorgundkaupangen博尔贡凯于庞恩 feud.Barony = &博尔贡凯于庞恩BorgundkaupangenBarony{}

func init() {
    f := BBorgundkaupangen博尔贡凯于庞恩.(*博尔贡凯于庞恩BorgundkaupangenBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "borgundkaupangen",
		TitleName: "博尔贡凯于庞恩",
		TitleCode: "b_borgundkaupangen",
	}
}
