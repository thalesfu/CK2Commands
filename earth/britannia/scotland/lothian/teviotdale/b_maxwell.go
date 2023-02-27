package teviotdale

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 马克斯韦尔MaxwellBarony struct {
	feud.BaseBarony
}

var BMaxwell马克斯韦尔 feud.Barony = &马克斯韦尔MaxwellBarony{}

func init() {
    f := BMaxwell马克斯韦尔.(*马克斯韦尔MaxwellBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "maxwell",
		TitleName: "马克斯韦尔",
		TitleCode: "b_maxwell",
	}
}
