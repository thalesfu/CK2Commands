package tana

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 巴泰斯克BatayskBarony struct {
	feud.BaseBarony
}

var BBataysk巴泰斯克 feud.Barony = &巴泰斯克BatayskBarony{}

func init() {
    f := BBataysk巴泰斯克.(*巴泰斯克BatayskBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "bataysk",
		TitleName: "巴泰斯克",
		TitleCode: "b_bataysk",
	}
}
