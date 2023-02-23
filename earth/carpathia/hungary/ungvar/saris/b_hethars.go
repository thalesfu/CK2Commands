package saris

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 海塔尔什HetharsBarony struct {
	feud.BaseBarony
}

var BHethars海塔尔什 feud.Barony = &海塔尔什HetharsBarony{}

func init() {
	f := BHethars海塔尔什.(*海塔尔什HetharsBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "hethars",
		TitleName: "海塔尔什",
		TitleCode: "b_hethars",
	}
}
