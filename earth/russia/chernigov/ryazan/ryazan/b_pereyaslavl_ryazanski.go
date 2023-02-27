package ryazan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 佩列亚斯拉夫尔梁赞斯基Pereyaslavl_ryazanskiBarony struct {
	feud.BaseBarony
}

var BPereyaslavl_ryazanski佩列亚斯拉夫尔梁赞斯基 feud.Barony = &佩列亚斯拉夫尔梁赞斯基Pereyaslavl_ryazanskiBarony{}

func init() {
    f := BPereyaslavl_ryazanski佩列亚斯拉夫尔梁赞斯基.(*佩列亚斯拉夫尔梁赞斯基Pereyaslavl_ryazanskiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "pereyaslavl_ryazanski",
		TitleName: "佩列亚斯拉夫尔梁赞斯基",
		TitleCode: "b_pereyaslavl_ryazanski",
	}
}
