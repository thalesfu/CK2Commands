package sarasvata_mandala

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 摩诃罗ModheraBarony struct {
	feud.BaseBarony
}

var BModhera摩诃罗 feud.Barony = &摩诃罗ModheraBarony{}

func init() {
    f := BModhera摩诃罗.(*摩诃罗ModheraBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "modhera",
		TitleName: "摩诃罗",
		TitleCode: "b_modhera",
	}
}
