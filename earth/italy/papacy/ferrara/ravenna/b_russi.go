package ravenna

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 鲁西RussiBarony struct {
	feud.BaseBarony
}

var BRussi鲁西 feud.Barony = &鲁西RussiBarony{}

func init() {
	f := BRussi鲁西.(*鲁西RussiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "russi",
		TitleName: "鲁西",
		TitleCode: "b_russi",
	}
}
