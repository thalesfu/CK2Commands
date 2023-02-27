package fyn

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 米泽尔法特MiddelfartBarony struct {
	feud.BaseBarony
}

var BMiddelfart米泽尔法特 feud.Barony = &米泽尔法特MiddelfartBarony{}

func init() {
    f := BMiddelfart米泽尔法特.(*米泽尔法特MiddelfartBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "middelfart",
		TitleName: "米泽尔法特",
		TitleCode: "b_middelfart",
	}
}
