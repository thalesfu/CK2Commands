package ikh_bogd

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 克烈KhereidBarony struct {
	feud.BaseBarony
}

var BKhereid克烈 feud.Barony = &克烈KhereidBarony{}

func init() {
    f := BKhereid克烈.(*克烈KhereidBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "khereid",
		TitleName: "克烈",
		TitleCode: "b_khereid",
	}
}
