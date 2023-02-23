package tikhvine

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 季赫温TikhvineBarony struct {
	feud.BaseBarony
}

var BTikhvine季赫温 feud.Barony = &季赫温TikhvineBarony{}

func init() {
	f := BTikhvine季赫温.(*季赫温TikhvineBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "tikhvine",
		TitleName: "季赫温",
		TitleCode: "b_tikhvine",
	}
}
