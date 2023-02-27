package yuryev

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 巴夫列内BavlenyBarony struct {
	feud.BaseBarony
}

var BBavleny巴夫列内 feud.Barony = &巴夫列内BavlenyBarony{}

func init() {
    f := BBavleny巴夫列内.(*巴夫列内BavlenyBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "bavleny",
		TitleName: "巴夫列内",
		TitleCode: "b_bavleny",
	}
}
