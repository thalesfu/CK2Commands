package ilam

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 察哈尔塔赫ChahartaghiBarony struct {
	feud.BaseBarony
}

var BChahartaghi察哈尔塔赫 feud.Barony = &察哈尔塔赫ChahartaghiBarony{}

func init() {
	f := BChahartaghi察哈尔塔赫.(*察哈尔塔赫ChahartaghiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "chahartaghi",
		TitleName: "察哈尔塔赫",
		TitleCode: "b_chahartaghi",
	}
}
