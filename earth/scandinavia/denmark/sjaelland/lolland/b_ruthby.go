package lolland

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 勒兹比RuthbyBarony struct {
	feud.BaseBarony
}

var BRuthby勒兹比 feud.Barony = &勒兹比RuthbyBarony{}

func init() {
	f := BRuthby勒兹比.(*勒兹比RuthbyBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ruthby",
		TitleName: "勒兹比",
		TitleCode: "b_ruthby",
	}
}
