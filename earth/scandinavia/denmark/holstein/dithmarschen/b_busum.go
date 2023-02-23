package dithmarschen

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 比苏姆BusumBarony struct {
	feud.BaseBarony
}

var BBusum比苏姆 feud.Barony = &比苏姆BusumBarony{}

func init() {
	f := BBusum比苏姆.(*比苏姆BusumBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "busum",
		TitleName: "比苏姆",
		TitleCode: "b_busum",
	}
}
