package orkney

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 伯赛BirsayBarony struct {
	feud.BaseBarony
}

var BBirsay伯赛 feud.Barony = &伯赛BirsayBarony{}

func init() {
	f := BBirsay伯赛.(*伯赛BirsayBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "birsay",
		TitleName: "伯赛",
		TitleCode: "b_birsay",
	}
}
