package treviso

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿索拉AsolaBarony struct {
	feud.BaseBarony
}

var BAsola阿索拉 feud.Barony = &阿索拉AsolaBarony{}

func init() {
	f := BAsola阿索拉.(*阿索拉AsolaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "asola",
		TitleName: "阿索拉",
		TitleCode: "b_asola",
	}
}
