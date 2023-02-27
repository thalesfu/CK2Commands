package druz

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 塞涅里赫ShannirehBarony struct {
	feud.BaseBarony
}

var BShannireh塞涅里赫 feud.Barony = &塞涅里赫ShannirehBarony{}

func init() {
    f := BShannireh塞涅里赫.(*塞涅里赫ShannirehBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "shannireh",
		TitleName: "塞涅里赫",
		TitleCode: "b_shannireh",
	}
}
