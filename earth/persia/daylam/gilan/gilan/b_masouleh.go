package gilan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 马苏莱MasoulehBarony struct {
	feud.BaseBarony
}

var BMasouleh马苏莱 feud.Barony = &马苏莱MasoulehBarony{}

func init() {
	f := BMasouleh马苏莱.(*马苏莱MasoulehBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "masouleh",
		TitleName: "马苏莱",
		TitleCode: "b_masouleh",
	}
}
