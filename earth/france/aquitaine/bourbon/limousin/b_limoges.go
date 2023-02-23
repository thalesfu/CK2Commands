package limousin

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 利摩日LimogesBarony struct {
	feud.BaseBarony
}

var BLimoges利摩日 feud.Barony = &利摩日LimogesBarony{}

func init() {
	f := BLimoges利摩日.(*利摩日LimogesBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "limoges",
		TitleName: "利摩日",
		TitleCode: "b_limoges",
	}
}
