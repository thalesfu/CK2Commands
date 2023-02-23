package austisland

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 雷扎尔菲亚尔ReydarfjallBarony struct {
	feud.BaseBarony
}

var BReydarfjall雷扎尔菲亚尔 feud.Barony = &雷扎尔菲亚尔ReydarfjallBarony{}

func init() {
	f := BReydarfjall雷扎尔菲亚尔.(*雷扎尔菲亚尔ReydarfjallBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "reydarfjall",
		TitleName: "雷扎尔菲亚尔",
		TitleCode: "b_reydarfjall",
	}
}
