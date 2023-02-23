package rohana

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 摩诃伽摩MahagamaBarony struct {
	feud.BaseBarony
}

var BMahagama摩诃伽摩 feud.Barony = &摩诃伽摩MahagamaBarony{}

func init() {
	f := BMahagama摩诃伽摩.(*摩诃伽摩MahagamaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "mahagama",
		TitleName: "摩诃伽摩",
		TitleCode: "b_mahagama",
	}
}
