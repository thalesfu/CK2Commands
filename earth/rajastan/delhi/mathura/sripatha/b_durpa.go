package sripatha

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 补波DurpaBarony struct {
	feud.BaseBarony
}

var BDurpa补波 feud.Barony = &补波DurpaBarony{}

func init() {
    f := BDurpa补波.(*补波DurpaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "durpa",
		TitleName: "补波",
		TitleCode: "b_durpa",
	}
}
