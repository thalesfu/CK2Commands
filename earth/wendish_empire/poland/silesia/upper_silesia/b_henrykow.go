package upper_silesia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 亨利库夫HenrykowBarony struct {
	feud.BaseBarony
}

var BHenrykow亨利库夫 feud.Barony = &亨利库夫HenrykowBarony{}

func init() {
    f := BHenrykow亨利库夫.(*亨利库夫HenrykowBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "henrykow",
		TitleName: "亨利库夫",
		TitleCode: "b_henrykow",
	}
}
