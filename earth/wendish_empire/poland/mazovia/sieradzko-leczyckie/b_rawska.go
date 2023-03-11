package sieradzko-leczyckie

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 拉瓦RawskaBarony struct {
	feud.BaseBarony
}

var BRawska拉瓦 feud.Barony = &拉瓦RawskaBarony{}

func init() {
    f := BRawska拉瓦.(*拉瓦RawskaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "rawska",
		TitleName: "拉瓦",
		TitleCode: "b_rawska",
	}
}
