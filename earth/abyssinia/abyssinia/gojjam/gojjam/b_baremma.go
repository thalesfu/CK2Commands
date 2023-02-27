package gojjam

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 巴雷马BaremmaBarony struct {
	feud.BaseBarony
}

var BBaremma巴雷马 feud.Barony = &巴雷马BaremmaBarony{}

func init() {
    f := BBaremma巴雷马.(*巴雷马BaremmaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "baremma",
		TitleName: "巴雷马",
		TitleCode: "b_baremma",
	}
}
