package kufa

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 布赛亚BussayyahBarony struct {
	feud.BaseBarony
}

var BBussayyah布赛亚 feud.Barony = &布赛亚BussayyahBarony{}

func init() {
	f := BBussayyah布赛亚.(*布赛亚BussayyahBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "bussayyah",
		TitleName: "布赛亚",
		TitleCode: "b_bussayyah",
	}
}
