package ranikot

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 旬陀JandarBarony struct {
	feud.BaseBarony
}

var BJandar旬陀 feud.Barony = &旬陀JandarBarony{}

func init() {
    f := BJandar旬陀.(*旬陀JandarBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "jandar",
		TitleName: "旬陀",
		TitleCode: "b_jandar",
	}
}
