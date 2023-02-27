package ranikot

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿布AbuBarony struct {
	feud.BaseBarony
}

var BAbu阿布 feud.Barony = &阿布AbuBarony{}

func init() {
    f := BAbu阿布.(*阿布AbuBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "abu",
		TitleName: "阿布",
		TitleCode: "b_abu",
	}
}
