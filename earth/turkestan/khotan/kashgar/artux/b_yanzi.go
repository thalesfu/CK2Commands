package artux

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 岩姿YanziBarony struct {
	feud.BaseBarony
}

var BYanzi岩姿 feud.Barony = &岩姿YanziBarony{}

func init() {
    f := BYanzi岩姿.(*岩姿YanziBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "yanzi",
		TitleName: "岩姿",
		TitleCode: "b_yanzi",
	}
}
