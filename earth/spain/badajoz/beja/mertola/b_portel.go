package mertola

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 波特尔PortelBarony struct {
	feud.BaseBarony
}

var BPortel波特尔 feud.Barony = &波特尔PortelBarony{}

func init() {
    f := BPortel波特尔.(*波特尔PortelBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "portel",
		TitleName: "波特尔",
		TitleCode: "b_portel",
	}
}
