package cherchen

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 左末ZuomoBarony struct {
	feud.BaseBarony
}

var BZuomo左末 feud.Barony = &左末ZuomoBarony{}

func init() {
	f := BZuomo左末.(*左末ZuomoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "zuomo",
		TitleName: "左末",
		TitleCode: "b_zuomo",
	}
}
