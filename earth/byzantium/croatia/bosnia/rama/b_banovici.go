package rama

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 巴诺维契BanoviciBarony struct {
	feud.BaseBarony
}

var BBanovici巴诺维契 feud.Barony = &巴诺维契BanoviciBarony{}

func init() {
    f := BBanovici巴诺维契.(*巴诺维契BanoviciBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "banovici",
		TitleName: "巴诺维契",
		TitleCode: "b_banovici",
	}
}
