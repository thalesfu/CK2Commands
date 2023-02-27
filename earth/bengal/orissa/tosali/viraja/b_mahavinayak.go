package viraja

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 摩诃毗那夜迦MahavinayakBarony struct {
	feud.BaseBarony
}

var BMahavinayak摩诃毗那夜迦 feud.Barony = &摩诃毗那夜迦MahavinayakBarony{}

func init() {
    f := BMahavinayak摩诃毗那夜迦.(*摩诃毗那夜迦MahavinayakBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "mahavinayak",
		TitleName: "摩诃毗那夜迦",
		TitleCode: "b_mahavinayak",
	}
}
