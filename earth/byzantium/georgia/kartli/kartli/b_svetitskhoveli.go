package kartli

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 斯韦季茨霍韦利SvetitskhoveliBarony struct {
	feud.BaseBarony
}

var BSvetitskhoveli斯韦季茨霍韦利 feud.Barony = &斯韦季茨霍韦利SvetitskhoveliBarony{}

func init() {
    f := BSvetitskhoveli斯韦季茨霍韦利.(*斯韦季茨霍韦利SvetitskhoveliBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "svetitskhoveli",
		TitleName: "斯韦季茨霍韦利",
		TitleCode: "b_svetitskhoveli",
	}
}
