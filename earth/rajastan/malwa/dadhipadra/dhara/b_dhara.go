package dhara

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 陀罗DharaBarony struct {
	feud.BaseBarony
}

var BDhara陀罗 feud.Barony = &陀罗DharaBarony{}

func init() {
    f := BDhara陀罗.(*陀罗DharaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "dhara",
		TitleName: "陀罗",
		TitleCode: "b_dhara",
	}
}
