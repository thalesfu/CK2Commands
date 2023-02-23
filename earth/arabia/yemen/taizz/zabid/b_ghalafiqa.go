package zabid

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 盖拉菲加GhalafiqaBarony struct {
	feud.BaseBarony
}

var BGhalafiqa盖拉菲加 feud.Barony = &盖拉菲加GhalafiqaBarony{}

func init() {
	f := BGhalafiqa盖拉菲加.(*盖拉菲加GhalafiqaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ghalafiqa",
		TitleName: "盖拉菲加",
		TitleCode: "b_ghalafiqa",
	}
}
