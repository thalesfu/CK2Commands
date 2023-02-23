package chauragarh

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 菩诃尼BohaniBarony struct {
	feud.BaseBarony
}

var BBohani菩诃尼 feud.Barony = &菩诃尼BohaniBarony{}

func init() {
	f := BBohani菩诃尼.(*菩诃尼BohaniBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "bohani",
		TitleName: "菩诃尼",
		TitleCode: "b_bohani",
	}
}
