package dipalpur

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 树俱般陀JongbandhaBarony struct {
	feud.BaseBarony
}

var BJongbandha树俱般陀 feud.Barony = &树俱般陀JongbandhaBarony{}

func init() {
    f := BJongbandha树俱般陀.(*树俱般陀JongbandhaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "jongbandha",
		TitleName: "树俱般陀",
		TitleCode: "b_jongbandha",
	}
}
