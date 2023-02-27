package begemder

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 但雅DanyaBarony struct {
	feud.BaseBarony
}

var BDanya但雅 feud.Barony = &但雅DanyaBarony{}

func init() {
    f := BDanya但雅.(*但雅DanyaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "danya",
		TitleName: "但雅",
		TitleCode: "b_danya",
	}
}
