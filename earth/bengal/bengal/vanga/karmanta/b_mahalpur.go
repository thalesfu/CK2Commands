package karmanta

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 摩诃罗补罗MahalpurBarony struct {
	feud.BaseBarony
}

var BMahalpur摩诃罗补罗 feud.Barony = &摩诃罗补罗MahalpurBarony{}

func init() {
	f := BMahalpur摩诃罗补罗.(*摩诃罗补罗MahalpurBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "mahalpur",
		TitleName: "摩诃罗补罗",
		TitleCode: "b_mahalpur",
	}
}
