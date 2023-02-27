package zatec

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 比利纳BilinaBarony struct {
	feud.BaseBarony
}

var BBilina比利纳 feud.Barony = &比利纳BilinaBarony{}

func init() {
    f := BBilina比利纳.(*比利纳BilinaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "bilina",
		TitleName: "比利纳",
		TitleCode: "b_bilina",
	}
}
