package yamalia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 纳济姆NazymBarony struct {
	feud.BaseBarony
}

var BNazym纳济姆 feud.Barony = &纳济姆NazymBarony{}

func init() {
    f := BNazym纳济姆.(*纳济姆NazymBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "nazym",
		TitleName: "纳济姆",
		TitleCode: "b_nazym",
	}
}
