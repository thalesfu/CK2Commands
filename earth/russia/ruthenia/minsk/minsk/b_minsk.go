package minsk

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 明斯克MinskBarony struct {
	feud.BaseBarony
}

var BMinsk明斯克 feud.Barony = &明斯克MinskBarony{}

func init() {
	f := BMinsk明斯克.(*明斯克MinskBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "minsk",
		TitleName: "明斯克",
		TitleCode: "b_minsk",
	}
}
