package strathearn

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 因彻法莱InchaffrayBarony struct {
	feud.BaseBarony
}

var BInchaffray因彻法莱 feud.Barony = &因彻法莱InchaffrayBarony{}

func init() {
    f := BInchaffray因彻法莱.(*因彻法莱InchaffrayBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "inchaffray",
		TitleName: "因彻法莱",
		TitleCode: "b_inchaffray",
	}
}
