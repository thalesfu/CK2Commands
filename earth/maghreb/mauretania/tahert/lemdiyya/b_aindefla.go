package lemdiyya

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 艾因德夫拉AindeflaBarony struct {
	feud.BaseBarony
}

var BAindefla艾因德夫拉 feud.Barony = &艾因德夫拉AindeflaBarony{}

func init() {
    f := BAindefla艾因德夫拉.(*艾因德夫拉AindeflaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "aindefla",
		TitleName: "艾因德夫拉",
		TitleCode: "b_aindefla",
	}
}
