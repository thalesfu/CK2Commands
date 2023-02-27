package pundravardhana

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 摩诃萨傥那姞利呬MahasthangarhBarony struct {
	feud.BaseBarony
}

var BMahasthangarh摩诃萨傥那姞利呬 feud.Barony = &摩诃萨傥那姞利呬MahasthangarhBarony{}

func init() {
    f := BMahasthangarh摩诃萨傥那姞利呬.(*摩诃萨傥那姞利呬MahasthangarhBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "mahasthangarh",
		TitleName: "摩诃萨傥那姞利呬",
		TitleCode: "b_mahasthangarh",
	}
}
