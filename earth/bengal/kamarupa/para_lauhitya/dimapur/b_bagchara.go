package dimapur

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 薄伽遮罗BagcharaBarony struct {
	feud.BaseBarony
}

var BBagchara薄伽遮罗 feud.Barony = &薄伽遮罗BagcharaBarony{}

func init() {
    f := BBagchara薄伽遮罗.(*薄伽遮罗BagcharaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "bagchara",
		TitleName: "薄伽遮罗",
		TitleCode: "b_bagchara",
	}
}
