package vodi

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 诺索克NosokBarony struct {
	feud.BaseBarony
}

var BNosok诺索克 feud.Barony = &诺索克NosokBarony{}

func init() {
    f := BNosok诺索克.(*诺索克NosokBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "nosok",
		TitleName: "诺索克",
		TitleCode: "b_nosok",
	}
}
