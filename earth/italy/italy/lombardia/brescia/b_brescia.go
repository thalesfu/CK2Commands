package brescia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 布雷西亚BresciaBarony struct {
	feud.BaseBarony
}

var BBrescia布雷西亚 feud.Barony = &布雷西亚BresciaBarony{}

func init() {
	f := BBrescia布雷西亚.(*布雷西亚BresciaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "brescia",
		TitleName: "布雷西亚",
		TitleCode: "b_brescia",
	}
}
