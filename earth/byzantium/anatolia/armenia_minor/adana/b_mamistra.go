package adana

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 曼密西MamistraBarony struct {
	feud.BaseBarony
}

var BMamistra曼密西 feud.Barony = &曼密西MamistraBarony{}

func init() {
    f := BMamistra曼密西.(*曼密西MamistraBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "mamistra",
		TitleName: "曼密西",
		TitleCode: "b_mamistra",
	}
}
