package aland

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 哥德比GodbyBarony struct {
	feud.BaseBarony
}

var BGodby哥德比 feud.Barony = &哥德比GodbyBarony{}

func init() {
    f := BGodby哥德比.(*哥德比GodbyBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "godby",
		TitleName: "哥德比",
		TitleCode: "b_godby",
	}
}
