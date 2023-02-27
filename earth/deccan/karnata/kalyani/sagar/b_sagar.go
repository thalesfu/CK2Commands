package sagar

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 娑伽罗SagarBarony struct {
	feud.BaseBarony
}

var BSagar娑伽罗 feud.Barony = &娑伽罗SagarBarony{}

func init() {
    f := BSagar娑伽罗.(*娑伽罗SagarBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "sagar",
		TitleName: "娑伽罗",
		TitleCode: "b_sagar",
	}
}
