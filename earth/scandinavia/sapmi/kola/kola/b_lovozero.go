package kola

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 洛沃泽罗LovozeroBarony struct {
	feud.BaseBarony
}

var BLovozero洛沃泽罗 feud.Barony = &洛沃泽罗LovozeroBarony{}

func init() {
    f := BLovozero洛沃泽罗.(*洛沃泽罗LovozeroBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "lovozero",
		TitleName: "洛沃泽罗",
		TitleCode: "b_lovozero",
	}
}
