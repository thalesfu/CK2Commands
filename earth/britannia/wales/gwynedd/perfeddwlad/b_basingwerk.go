package perfeddwlad

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 巴辛韦克BasingwerkBarony struct {
	feud.BaseBarony
}

var BBasingwerk巴辛韦克 feud.Barony = &巴辛韦克BasingwerkBarony{}

func init() {
	f := BBasingwerk巴辛韦克.(*巴辛韦克BasingwerkBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "basingwerk",
		TitleName: "巴辛韦克",
		TitleCode: "b_basingwerk",
	}
}
