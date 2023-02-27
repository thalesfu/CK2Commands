package maldives

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 摩诃罗MahalBarony struct {
	feud.BaseBarony
}

var BMahal摩诃罗 feud.Barony = &摩诃罗MahalBarony{}

func init() {
    f := BMahal摩诃罗.(*摩诃罗MahalBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "mahal",
		TitleName: "摩诃罗",
		TitleCode: "b_mahal",
	}
}
