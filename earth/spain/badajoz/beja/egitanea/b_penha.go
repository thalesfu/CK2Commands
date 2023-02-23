package egitanea

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 佩尼亚PenhaBarony struct {
	feud.BaseBarony
}

var BPenha佩尼亚 feud.Barony = &佩尼亚PenhaBarony{}

func init() {
	f := BPenha佩尼亚.(*佩尼亚PenhaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "penha",
		TitleName: "佩尼亚",
		TitleCode: "b_penha",
	}
}
