package calarasi

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 科尔纳泰CornateiBarony struct {
	feud.BaseBarony
}

var BCornatei科尔纳泰 feud.Barony = &科尔纳泰CornateiBarony{}

func init() {
    f := BCornatei科尔纳泰.(*科尔纳泰CornateiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "cornatei",
		TitleName: "科尔纳泰",
		TitleCode: "b_cornatei",
	}
}
