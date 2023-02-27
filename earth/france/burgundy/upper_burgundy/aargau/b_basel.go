package aargau

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 巴塞尔BaselBarony struct {
	feud.BaseBarony
}

var BBasel巴塞尔 feud.Barony = &巴塞尔BaselBarony{}

func init() {
    f := BBasel巴塞尔.(*巴塞尔BaselBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "basel",
		TitleName: "巴塞尔",
		TitleCode: "b_basel",
	}
}
