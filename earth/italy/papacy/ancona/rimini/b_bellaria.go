package rimini

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 贝拉里亚BellariaBarony struct {
	feud.BaseBarony
}

var BBellaria贝拉里亚 feud.Barony = &贝拉里亚BellariaBarony{}

func init() {
    f := BBellaria贝拉里亚.(*贝拉里亚BellariaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "bellaria",
		TitleName: "贝拉里亚",
		TitleCode: "b_bellaria",
	}
}
