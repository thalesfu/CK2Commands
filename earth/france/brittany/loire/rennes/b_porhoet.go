package rennes

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 波尔厄PorhoetBarony struct {
	feud.BaseBarony
}

var BPorhoet波尔厄 feud.Barony = &波尔厄PorhoetBarony{}

func init() {
    f := BPorhoet波尔厄.(*波尔厄PorhoetBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "porhoet",
		TitleName: "波尔厄",
		TitleCode: "b_porhoet",
	}
}
