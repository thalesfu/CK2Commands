package kolguyev

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 佩尔瓦亚PervayaBarony struct {
	feud.BaseBarony
}

var BPervaya佩尔瓦亚 feud.Barony = &佩尔瓦亚PervayaBarony{}

func init() {
    f := BPervaya佩尔瓦亚.(*佩尔瓦亚PervayaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "pervaya",
		TitleName: "佩尔瓦亚",
		TitleCode: "b_pervaya",
	}
}
