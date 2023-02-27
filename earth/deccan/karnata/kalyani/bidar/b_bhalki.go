package bidar

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 婆吉BhalkiBarony struct {
	feud.BaseBarony
}

var BBhalki婆吉 feud.Barony = &婆吉BhalkiBarony{}

func init() {
    f := BBhalki婆吉.(*婆吉BhalkiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "bhalki",
		TitleName: "婆吉",
		TitleCode: "b_bhalki",
	}
}
