package taron

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿赫拉特AkhlatBarony struct {
	feud.BaseBarony
}

var BAkhlat阿赫拉特 feud.Barony = &阿赫拉特AkhlatBarony{}

func init() {
	f := BAkhlat阿赫拉特.(*阿赫拉特AkhlatBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "akhlat",
		TitleName: "阿赫拉特",
		TitleCode: "b_akhlat",
	}
}
