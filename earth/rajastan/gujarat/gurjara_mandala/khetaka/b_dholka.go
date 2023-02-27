package khetaka

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 陀尔迦DholkaBarony struct {
	feud.BaseBarony
}

var BDholka陀尔迦 feud.Barony = &陀尔迦DholkaBarony{}

func init() {
    f := BDholka陀尔迦.(*陀尔迦DholkaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "dholka",
		TitleName: "陀尔迦",
		TitleCode: "b_dholka",
	}
}
