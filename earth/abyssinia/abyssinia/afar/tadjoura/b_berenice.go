package tadjoura

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 贝勒尼基BereniceBarony struct {
	feud.BaseBarony
}

var BBerenice贝勒尼基 feud.Barony = &贝勒尼基BereniceBarony{}

func init() {
    f := BBerenice贝勒尼基.(*贝勒尼基BereniceBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "berenice",
		TitleName: "贝勒尼基",
		TitleCode: "b_berenice",
	}
}
