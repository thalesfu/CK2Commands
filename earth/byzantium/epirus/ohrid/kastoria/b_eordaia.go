package kastoria

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 欧耳代亚EordaiaBarony struct {
	feud.BaseBarony
}

var BEordaia欧耳代亚 feud.Barony = &欧耳代亚EordaiaBarony{}

func init() {
    f := BEordaia欧耳代亚.(*欧耳代亚EordaiaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "eordaia",
		TitleName: "欧耳代亚",
		TitleCode: "b_eordaia",
	}
}
