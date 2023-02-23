package pereyaslavl

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 佩列亚斯拉夫尔PereyaslavlBarony struct {
	feud.BaseBarony
}

var BPereyaslavl佩列亚斯拉夫尔 feud.Barony = &佩列亚斯拉夫尔PereyaslavlBarony{}

func init() {
	f := BPereyaslavl佩列亚斯拉夫尔.(*佩列亚斯拉夫尔PereyaslavlBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "pereyaslavl",
		TitleName: "佩列亚斯拉夫尔",
		TitleCode: "b_pereyaslavl",
	}
}
