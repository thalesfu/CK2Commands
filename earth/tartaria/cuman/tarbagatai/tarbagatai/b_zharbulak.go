package tarbagatai

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 扎尔布拉克ZharbulakBarony struct {
	feud.BaseBarony
}

var BZharbulak扎尔布拉克 feud.Barony = &扎尔布拉克ZharbulakBarony{}

func init() {
	f := BZharbulak扎尔布拉克.(*扎尔布拉克ZharbulakBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "zharbulak",
		TitleName: "扎尔布拉克",
		TitleCode: "b_zharbulak",
	}
}
