package naupactus

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 纳夫帕克托斯NaupaktosBarony struct {
	feud.BaseBarony
}

var BNaupaktos纳夫帕克托斯 feud.Barony = &纳夫帕克托斯NaupaktosBarony{}

func init() {
	f := BNaupaktos纳夫帕克托斯.(*纳夫帕克托斯NaupaktosBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "naupaktos",
		TitleName: "纳夫帕克托斯",
		TitleCode: "b_naupaktos",
	}
}
