package surrey

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 彻特西ChertseyBarony struct {
	feud.BaseBarony
}

var BChertsey彻特西 feud.Barony = &彻特西ChertseyBarony{}

func init() {
    f := BChertsey彻特西.(*彻特西ChertseyBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "chertsey",
		TitleName: "彻特西",
		TitleCode: "b_chertsey",
	}
}
