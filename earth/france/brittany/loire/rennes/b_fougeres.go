package rennes

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 富热尔FougeresBarony struct {
	feud.BaseBarony
}

var BFougeres富热尔 feud.Barony = &富热尔FougeresBarony{}

func init() {
	f := BFougeres富热尔.(*富热尔FougeresBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "fougeres",
		TitleName: "富热尔",
		TitleCode: "b_fougeres",
	}
}
