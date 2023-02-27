package zurichgau

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 拉珀斯维尔RapperswilBarony struct {
	feud.BaseBarony
}

var BRapperswil拉珀斯维尔 feud.Barony = &拉珀斯维尔RapperswilBarony{}

func init() {
    f := BRapperswil拉珀斯维尔.(*拉珀斯维尔RapperswilBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "rapperswil",
		TitleName: "拉珀斯维尔",
		TitleCode: "b_rapperswil",
	}
}
