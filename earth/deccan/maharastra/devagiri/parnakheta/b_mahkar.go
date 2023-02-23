package parnakheta

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 摩诃羯罗MahkarBarony struct {
	feud.BaseBarony
}

var BMahkar摩诃羯罗 feud.Barony = &摩诃羯罗MahkarBarony{}

func init() {
	f := BMahkar摩诃羯罗.(*摩诃羯罗MahkarBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "mahkar",
		TitleName: "摩诃羯罗",
		TitleCode: "b_mahkar",
	}
}
