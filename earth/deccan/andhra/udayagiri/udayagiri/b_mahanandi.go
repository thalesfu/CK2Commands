package udayagiri

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 摩诃难提MahanandiBarony struct {
	feud.BaseBarony
}

var BMahanandi摩诃难提 feud.Barony = &摩诃难提MahanandiBarony{}

func init() {
	f := BMahanandi摩诃难提.(*摩诃难提MahanandiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "mahanandi",
		TitleName: "摩诃难提",
		TitleCode: "b_mahanandi",
	}
}
