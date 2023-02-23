package belgorod

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 别尔哥罗德BelgorodBarony struct {
	feud.BaseBarony
}

var BBelgorod别尔哥罗德 feud.Barony = &别尔哥罗德BelgorodBarony{}

func init() {
	f := BBelgorod别尔哥罗德.(*别尔哥罗德BelgorodBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "belgorod",
		TitleName: "别尔哥罗德",
		TitleCode: "b_belgorod",
	}
}
