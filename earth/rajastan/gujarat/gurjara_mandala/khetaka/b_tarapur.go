package khetaka

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 多罗城TarapurBarony struct {
	feud.BaseBarony
}

var BTarapur多罗城 feud.Barony = &多罗城TarapurBarony{}

func init() {
    f := BTarapur多罗城.(*多罗城TarapurBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "tarapur",
		TitleName: "多罗城",
		TitleCode: "b_tarapur",
	}
}
