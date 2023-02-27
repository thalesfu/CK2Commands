package pereyaslavl_zalessky

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 佩列亚斯拉夫尔PereyaslavlzalesskyBarony struct {
	feud.BaseBarony
}

var BPereyaslavlzalessky佩列亚斯拉夫尔 feud.Barony = &佩列亚斯拉夫尔PereyaslavlzalesskyBarony{}

func init() {
    f := BPereyaslavlzalessky佩列亚斯拉夫尔.(*佩列亚斯拉夫尔PereyaslavlzalesskyBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "pereyaslavlzalessky",
		TitleName: "佩列亚斯拉夫尔",
		TitleCode: "b_pereyaslavlzalessky",
	}
}
