package chlisselbourg

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 诺索克NossokBarony struct {
	feud.BaseBarony
}

var BNossok诺索克 feud.Barony = &诺索克NossokBarony{}

func init() {
    f := BNossok诺索克.(*诺索克NossokBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "nossok",
		TitleName: "诺索克",
		TitleCode: "b_nossok",
	}
}
