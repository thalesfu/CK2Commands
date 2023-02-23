package ivrea

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 科萨托CossatoBarony struct {
	feud.BaseBarony
}

var BCossato科萨托 feud.Barony = &科萨托CossatoBarony{}

func init() {
	f := BCossato科萨托.(*科萨托CossatoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "cossato",
		TitleName: "科萨托",
		TitleCode: "b_cossato",
	}
}
