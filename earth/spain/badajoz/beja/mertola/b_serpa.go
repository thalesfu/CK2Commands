package mertola

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 塞尔帕SerpaBarony struct {
	feud.BaseBarony
}

var BSerpa塞尔帕 feud.Barony = &塞尔帕SerpaBarony{}

func init() {
	f := BSerpa塞尔帕.(*塞尔帕SerpaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "serpa",
		TitleName: "塞尔帕",
		TitleCode: "b_serpa",
	}
}
