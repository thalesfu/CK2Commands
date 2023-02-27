package sagar

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 巴加瓦迪BagavadiBarony struct {
	feud.BaseBarony
}

var BBagavadi巴加瓦迪 feud.Barony = &巴加瓦迪BagavadiBarony{}

func init() {
    f := BBagavadi巴加瓦迪.(*巴加瓦迪BagavadiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "bagavadi",
		TitleName: "巴加瓦迪",
		TitleCode: "b_bagavadi",
	}
}
