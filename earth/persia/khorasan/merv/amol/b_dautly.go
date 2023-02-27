package amol

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 多特利DautlyBarony struct {
	feud.BaseBarony
}

var BDautly多特利 feud.Barony = &多特利DautlyBarony{}

func init() {
    f := BDautly多特利.(*多特利DautlyBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "dautly",
		TitleName: "多特利",
		TitleCode: "b_dautly",
	}
}
