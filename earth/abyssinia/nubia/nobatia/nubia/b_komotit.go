package nubia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 科莫蒂特KomotitBarony struct {
	feud.BaseBarony
}

var BKomotit科莫蒂特 feud.Barony = &科莫蒂特KomotitBarony{}

func init() {
	f := BKomotit科莫蒂特.(*科莫蒂特KomotitBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "komotit",
		TitleName: "科莫蒂特",
		TitleCode: "b_komotit",
	}
}
