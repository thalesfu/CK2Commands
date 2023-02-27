package ejin

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 马鬃山MazongshanBarony struct {
	feud.BaseBarony
}

var BMazongshan马鬃山 feud.Barony = &马鬃山MazongshanBarony{}

func init() {
    f := BMazongshan马鬃山.(*马鬃山MazongshanBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "mazongshan",
		TitleName: "马鬃山",
		TitleCode: "b_mazongshan",
	}
}
