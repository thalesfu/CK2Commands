package chelminskie

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 勒包LobauBarony struct {
	feud.BaseBarony
}

var BLobau勒包 feud.Barony = &勒包LobauBarony{}

func init() {
    f := BLobau勒包.(*勒包LobauBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "lobau",
		TitleName: "勒包",
		TitleCode: "b_lobau",
	}
}
