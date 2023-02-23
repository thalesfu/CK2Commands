package calarasi

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 费泰什蒂FetestiBarony struct {
	feud.BaseBarony
}

var BFetesti费泰什蒂 feud.Barony = &费泰什蒂FetestiBarony{}

func init() {
	f := BFetesti费泰什蒂.(*费泰什蒂FetestiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "fetesti",
		TitleName: "费泰什蒂",
		TitleCode: "b_fetesti",
	}
}
