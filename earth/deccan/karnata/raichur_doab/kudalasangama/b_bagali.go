package kudalasangama

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 巴加利BagaliBarony struct {
	feud.BaseBarony
}

var BBagali巴加利 feud.Barony = &巴加利BagaliBarony{}

func init() {
    f := BBagali巴加利.(*巴加利BagaliBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "bagali",
		TitleName: "巴加利",
		TitleCode: "b_bagali",
	}
}
