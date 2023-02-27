package bambuk

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 吉雅鲁GhiyaruBarony struct {
	feud.BaseBarony
}

var BGhiyaru吉雅鲁 feud.Barony = &吉雅鲁GhiyaruBarony{}

func init() {
    f := BGhiyaru吉雅鲁.(*吉雅鲁GhiyaruBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ghiyaru",
		TitleName: "吉雅鲁",
		TitleCode: "b_ghiyaru",
	}
}
