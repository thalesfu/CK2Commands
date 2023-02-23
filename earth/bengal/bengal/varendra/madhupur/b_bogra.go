package madhupur

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 菩迦罗BograBarony struct {
	feud.BaseBarony
}

var BBogra菩迦罗 feud.Barony = &菩迦罗BograBarony{}

func init() {
	f := BBogra菩迦罗.(*菩迦罗BograBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "bogra",
		TitleName: "菩迦罗",
		TitleCode: "b_bogra",
	}
}
