package lutsk

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 杜布诺DubnoBarony struct {
	feud.BaseBarony
}

var BDubno杜布诺 feud.Barony = &杜布诺DubnoBarony{}

func init() {
    f := BDubno杜布诺.(*杜布诺DubnoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "dubno",
		TitleName: "杜布诺",
		TitleCode: "b_dubno",
	}
}
