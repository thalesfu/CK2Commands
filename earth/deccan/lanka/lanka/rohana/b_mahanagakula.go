package rohana

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 摩诃那伽俱罗MahanagakulaBarony struct {
	feud.BaseBarony
}

var BMahanagakula摩诃那伽俱罗 feud.Barony = &摩诃那伽俱罗MahanagakulaBarony{}

func init() {
    f := BMahanagakula摩诃那伽俱罗.(*摩诃那伽俱罗MahanagakulaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "mahanagakula",
		TitleName: "摩诃那伽俱罗",
		TitleCode: "b_mahanagakula",
	}
}
