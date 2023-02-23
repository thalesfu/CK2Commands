package paphlagonia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 莱翁托波利斯LeontopolisBarony struct {
	feud.BaseBarony
}

var BLeontopolis莱翁托波利斯 feud.Barony = &莱翁托波利斯LeontopolisBarony{}

func init() {
	f := BLeontopolis莱翁托波利斯.(*莱翁托波利斯LeontopolisBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "leontopolis",
		TitleName: "莱翁托波利斯",
		TitleCode: "b_leontopolis",
	}
}
