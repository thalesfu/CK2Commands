package halogaland

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 勒因LeinBarony struct {
	feud.BaseBarony
}

var BLein勒因 feud.Barony = &勒因LeinBarony{}

func init() {
    f := BLein勒因.(*勒因LeinBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "lein",
		TitleName: "勒因",
		TitleCode: "b_lein",
	}
}
