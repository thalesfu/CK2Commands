package montbeliard

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 巴旺BavansBarony struct {
	feud.BaseBarony
}

var BBavans巴旺 feud.Barony = &巴旺BavansBarony{}

func init() {
    f := BBavans巴旺.(*巴旺BavansBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "bavans",
		TitleName: "巴旺",
		TitleCode: "b_bavans",
	}
}
