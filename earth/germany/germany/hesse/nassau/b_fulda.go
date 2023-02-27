package nassau

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 富尔达FuldaBarony struct {
	feud.BaseBarony
}

var BFulda富尔达 feud.Barony = &富尔达FuldaBarony{}

func init() {
    f := BFulda富尔达.(*富尔达FuldaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "fulda",
		TitleName: "富尔达",
		TitleCode: "b_fulda",
	}
}
