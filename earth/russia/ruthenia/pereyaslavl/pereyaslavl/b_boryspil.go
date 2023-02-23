package pereyaslavl

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 鲍里斯波尔BoryspilBarony struct {
	feud.BaseBarony
}

var BBoryspil鲍里斯波尔 feud.Barony = &鲍里斯波尔BoryspilBarony{}

func init() {
	f := BBoryspil鲍里斯波尔.(*鲍里斯波尔BoryspilBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "boryspil",
		TitleName: "鲍里斯波尔",
		TitleCode: "b_boryspil",
	}
}
