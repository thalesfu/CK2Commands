package delingha

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 河西HexiBarony struct {
	feud.BaseBarony
}

var BHexi河西 feud.Barony = &河西HexiBarony{}

func init() {
    f := BHexi河西.(*河西HexiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "hexi",
		TitleName: "河西",
		TitleCode: "b_hexi",
	}
}
