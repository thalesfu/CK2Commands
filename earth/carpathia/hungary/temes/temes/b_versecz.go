package temes

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 韦尔谢茨VerseczBarony struct {
	feud.BaseBarony
}

var BVersecz韦尔谢茨 feud.Barony = &韦尔谢茨VerseczBarony{}

func init() {
	f := BVersecz韦尔谢茨.(*韦尔谢茨VerseczBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "versecz",
		TitleName: "韦尔谢茨",
		TitleCode: "b_versecz",
	}
}
