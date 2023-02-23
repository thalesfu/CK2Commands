package bari

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 巴里BariBarony struct {
	feud.BaseBarony
}

var BBari巴里 feud.Barony = &巴里BariBarony{}

func init() {
	f := BBari巴里.(*巴里BariBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "bari",
		TitleName: "巴里",
		TitleCode: "b_bari",
	}
}
