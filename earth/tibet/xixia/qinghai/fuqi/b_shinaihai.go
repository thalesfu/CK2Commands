package fuqi

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 石乃亥ShinaihaiBarony struct {
	feud.BaseBarony
}

var BShinaihai石乃亥 feud.Barony = &石乃亥ShinaihaiBarony{}

func init() {
    f := BShinaihai石乃亥.(*石乃亥ShinaihaiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "shinaihai",
		TitleName: "石乃亥",
		TitleCode: "b_shinaihai",
	}
}
