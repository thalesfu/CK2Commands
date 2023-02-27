package dunois

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 马什努瓦MarchenoirBarony struct {
	feud.BaseBarony
}

var BMarchenoir马什努瓦 feud.Barony = &马什努瓦MarchenoirBarony{}

func init() {
    f := BMarchenoir马什努瓦.(*马什努瓦MarchenoirBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "marchenoir",
		TitleName: "马什努瓦",
		TitleCode: "b_marchenoir",
	}
}
