package phiti

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 摩诃迭他MahatitthaBarony struct {
	feud.BaseBarony
}

var BMahatittha摩诃迭他 feud.Barony = &摩诃迭他MahatitthaBarony{}

func init() {
    f := BMahatittha摩诃迭他.(*摩诃迭他MahatitthaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "mahatittha",
		TitleName: "摩诃迭他",
		TitleCode: "b_mahatittha",
	}
}
