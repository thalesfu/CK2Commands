package urzhar

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 曼扎尔ManzharBarony struct {
	feud.BaseBarony
}

var BManzhar曼扎尔 feud.Barony = &曼扎尔ManzharBarony{}

func init() {
	f := BManzhar曼扎尔.(*曼扎尔ManzharBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "manzhar",
		TitleName: "曼扎尔",
		TitleCode: "b_manzhar",
	}
}
