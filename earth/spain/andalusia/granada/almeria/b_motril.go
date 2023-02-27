package almeria

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 莫特里尔MotrilBarony struct {
	feud.BaseBarony
}

var BMotril莫特里尔 feud.Barony = &莫特里尔MotrilBarony{}

func init() {
    f := BMotril莫特里尔.(*莫特里尔MotrilBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "motril",
		TitleName: "莫特里尔",
		TitleCode: "b_motril",
	}
}
