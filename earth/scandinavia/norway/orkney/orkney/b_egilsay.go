package orkney

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 埃吉尔赛EgilsayBarony struct {
	feud.BaseBarony
}

var BEgilsay埃吉尔赛 feud.Barony = &埃吉尔赛EgilsayBarony{}

func init() {
    f := BEgilsay埃吉尔赛.(*埃吉尔赛EgilsayBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "egilsay",
		TitleName: "埃吉尔赛",
		TitleCode: "b_egilsay",
	}
}
