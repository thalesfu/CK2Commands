package esztergom

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 瑙吉格曼德NagyigmandBarony struct {
	feud.BaseBarony
}

var BNagyigmand瑙吉格曼德 feud.Barony = &瑙吉格曼德NagyigmandBarony{}

func init() {
	f := BNagyigmand瑙吉格曼德.(*瑙吉格曼德NagyigmandBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "nagyigmand",
		TitleName: "瑙吉格曼德",
		TitleCode: "b_nagyigmand",
	}
}
