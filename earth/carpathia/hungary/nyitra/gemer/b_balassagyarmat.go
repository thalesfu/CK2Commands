package gemer

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 包洛绍焦尔毛特BalassagyarmatBarony struct {
	feud.BaseBarony
}

var BBalassagyarmat包洛绍焦尔毛特 feud.Barony = &包洛绍焦尔毛特BalassagyarmatBarony{}

func init() {
    f := BBalassagyarmat包洛绍焦尔毛特.(*包洛绍焦尔毛特BalassagyarmatBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "balassagyarmat",
		TitleName: "包洛绍焦尔毛特",
		TitleCode: "b_balassagyarmat",
	}
}
