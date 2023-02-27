package marienburg

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 巴滕施泰因BartensteinBarony struct {
	feud.BaseBarony
}

var BBartenstein巴滕施泰因 feud.Barony = &巴滕施泰因BartensteinBarony{}

func init() {
    f := BBartenstein巴滕施泰因.(*巴滕施泰因BartensteinBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "bartenstein",
		TitleName: "巴滕施泰因",
		TitleCode: "b_bartenstein",
	}
}
