package kol

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 尼罗伐底NilavatiBarony struct {
	feud.BaseBarony
}

var BNilavati尼罗伐底 feud.Barony = &尼罗伐底NilavatiBarony{}

func init() {
	f := BNilavati尼罗伐底.(*尼罗伐底NilavatiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "nilavati",
		TitleName: "尼罗伐底",
		TitleCode: "b_nilavati",
	}
}
