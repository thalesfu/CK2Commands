package nalut

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿马姆里亚El_amarmriaBarony struct {
	feud.BaseBarony
}

var BEl_amarmria阿马姆里亚 feud.Barony = &阿马姆里亚El_amarmriaBarony{}

func init() {
    f := BEl_amarmria阿马姆里亚.(*阿马姆里亚El_amarmriaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "el_amarmria",
		TitleName: "阿马姆里亚",
		TitleCode: "b_el_amarmria",
	}
}
