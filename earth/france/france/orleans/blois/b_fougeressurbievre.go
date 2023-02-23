package blois

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 比耶夫尔河畔富热尔FougeressurbievreBarony struct {
	feud.BaseBarony
}

var BFougeressurbievre比耶夫尔河畔富热尔 feud.Barony = &比耶夫尔河畔富热尔FougeressurbievreBarony{}

func init() {
	f := BFougeressurbievre比耶夫尔河畔富热尔.(*比耶夫尔河畔富热尔FougeressurbievreBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "fougeressurbievre",
		TitleName: "比耶夫尔河畔富热尔",
		TitleCode: "b_fougeressurbievre",
	}
}
