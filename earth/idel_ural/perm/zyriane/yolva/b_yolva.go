package yolva

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 约尔瓦YolvaBarony struct {
	feud.BaseBarony
}

var BYolva约尔瓦 feud.Barony = &约尔瓦YolvaBarony{}

func init() {
    f := BYolva约尔瓦.(*约尔瓦YolvaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "yolva",
		TitleName: "约尔瓦",
		TitleCode: "b_yolva",
	}
}
