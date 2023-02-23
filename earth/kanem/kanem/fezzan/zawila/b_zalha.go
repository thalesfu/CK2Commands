package zawila

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 泽尔哈ZalhaBarony struct {
	feud.BaseBarony
}

var BZalha泽尔哈 feud.Barony = &泽尔哈ZalhaBarony{}

func init() {
	f := BZalha泽尔哈.(*泽尔哈ZalhaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "zalha",
		TitleName: "泽尔哈",
		TitleCode: "b_zalha",
	}
}
