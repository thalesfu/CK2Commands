package northampton

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 拉姆西RamseyBarony struct {
	feud.BaseBarony
}

var BRamsey拉姆西 feud.Barony = &拉姆西RamseyBarony{}

func init() {
    f := BRamsey拉姆西.(*拉姆西RamseyBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ramsey",
		TitleName: "拉姆西",
		TitleCode: "b_ramsey",
	}
}
