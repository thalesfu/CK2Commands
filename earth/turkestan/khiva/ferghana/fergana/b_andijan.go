package fergana

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 俺的干AndijanBarony struct {
	feud.BaseBarony
}

var BAndijan俺的干 feud.Barony = &俺的干AndijanBarony{}

func init() {
	f := BAndijan俺的干.(*俺的干AndijanBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "andijan",
		TitleName: "俺的干",
		TitleCode: "b_andijan",
	}
}
