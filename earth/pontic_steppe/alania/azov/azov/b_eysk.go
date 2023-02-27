package azov

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 叶伊斯克EyskBarony struct {
	feud.BaseBarony
}

var BEysk叶伊斯克 feud.Barony = &叶伊斯克EyskBarony{}

func init() {
    f := BEysk叶伊斯克.(*叶伊斯克EyskBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "eysk",
		TitleName: "叶伊斯克",
		TitleCode: "b_eysk",
	}
}
