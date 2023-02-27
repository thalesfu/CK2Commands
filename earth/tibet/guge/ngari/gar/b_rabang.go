package gar

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 热帮RabangBarony struct {
	feud.BaseBarony
}

var BRabang热帮 feud.Barony = &热帮RabangBarony{}

func init() {
    f := BRabang热帮.(*热帮RabangBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "rabang",
		TitleName: "热帮",
		TitleCode: "b_rabang",
	}
}
