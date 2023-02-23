package winchester

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 拉姆西RomseyBarony struct {
	feud.BaseBarony
}

var BRomsey拉姆西 feud.Barony = &拉姆西RomseyBarony{}

func init() {
	f := BRomsey拉姆西.(*拉姆西RomseyBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "romsey",
		TitleName: "拉姆西",
		TitleCode: "b_romsey",
	}
}
