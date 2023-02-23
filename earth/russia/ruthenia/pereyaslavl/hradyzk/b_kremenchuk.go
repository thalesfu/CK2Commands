package hradyzk

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 克列缅丘格KremenchukBarony struct {
	feud.BaseBarony
}

var BKremenchuk克列缅丘格 feud.Barony = &克列缅丘格KremenchukBarony{}

func init() {
	f := BKremenchuk克列缅丘格.(*克列缅丘格KremenchukBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kremenchuk",
		TitleName: "克列缅丘格",
		TitleCode: "b_kremenchuk",
	}
}
