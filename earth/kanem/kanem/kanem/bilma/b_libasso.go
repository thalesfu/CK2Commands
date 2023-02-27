package bilma

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 利巴索LibassoBarony struct {
	feud.BaseBarony
}

var BLibasso利巴索 feud.Barony = &利巴索LibassoBarony{}

func init() {
    f := BLibasso利巴索.(*利巴索LibassoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "libasso",
		TitleName: "利巴索",
		TitleCode: "b_libasso",
	}
}
