package urgell

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 巴利德沃伊ValledebohiBarony struct {
	feud.BaseBarony
}

var BValledebohi巴利德沃伊 feud.Barony = &巴利德沃伊ValledebohiBarony{}

func init() {
	f := BValledebohi巴利德沃伊.(*巴利德沃伊ValledebohiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "valledebohi",
		TitleName: "巴利德沃伊",
		TitleCode: "b_valledebohi",
	}
}
