package tjust

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 杨格尔斯鲁姆JangolsrumBarony struct {
	feud.BaseBarony
}

var BJangolsrum杨格尔斯鲁姆 feud.Barony = &杨格尔斯鲁姆JangolsrumBarony{}

func init() {
	f := BJangolsrum杨格尔斯鲁姆.(*杨格尔斯鲁姆JangolsrumBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "jangolsrum",
		TitleName: "杨格尔斯鲁姆",
		TitleCode: "b_jangolsrum",
	}
}
