package murzuk

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 迈尔祖格MurzukBarony struct {
	feud.BaseBarony
}

var BMurzuk迈尔祖格 feud.Barony = &迈尔祖格MurzukBarony{}

func init() {
    f := BMurzuk迈尔祖格.(*迈尔祖格MurzukBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "murzuk",
		TitleName: "迈尔祖格",
		TitleCode: "b_murzuk",
	}
}
