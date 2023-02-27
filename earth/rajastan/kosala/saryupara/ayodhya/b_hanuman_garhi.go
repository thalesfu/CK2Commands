package ayodhya

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 诃奴曼姞利呬Hanuman_garhiBarony struct {
	feud.BaseBarony
}

var BHanuman_garhi诃奴曼姞利呬 feud.Barony = &诃奴曼姞利呬Hanuman_garhiBarony{}

func init() {
    f := BHanuman_garhi诃奴曼姞利呬.(*诃奴曼姞利呬Hanuman_garhiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "hanuman_garhi",
		TitleName: "诃奴曼姞利呬",
		TitleCode: "b_hanuman_garhi",
	}
}
