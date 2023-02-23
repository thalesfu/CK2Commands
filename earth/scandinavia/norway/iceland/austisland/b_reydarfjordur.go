package austisland

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 雷扎尔菲厄泽ReydarfjordurBarony struct {
	feud.BaseBarony
}

var BReydarfjordur雷扎尔菲厄泽 feud.Barony = &雷扎尔菲厄泽ReydarfjordurBarony{}

func init() {
	f := BReydarfjordur雷扎尔菲厄泽.(*雷扎尔菲厄泽ReydarfjordurBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "reydarfjordur",
		TitleName: "雷扎尔菲厄泽",
		TitleCode: "b_reydarfjordur",
	}
}
