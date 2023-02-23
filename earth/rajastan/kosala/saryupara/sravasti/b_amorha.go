package sravasti

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿没诃AmorhaBarony struct {
	feud.BaseBarony
}

var BAmorha阿没诃 feud.Barony = &阿没诃AmorhaBarony{}

func init() {
	f := BAmorha阿没诃.(*阿没诃AmorhaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "amorha",
		TitleName: "阿没诃",
		TitleCode: "b_amorha",
	}
}
