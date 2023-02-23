package snassen

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 杰拉代JeradaBarony struct {
	feud.BaseBarony
}

var BJerada杰拉代 feud.Barony = &杰拉代JeradaBarony{}

func init() {
	f := BJerada杰拉代.(*杰拉代JeradaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "jerada",
		TitleName: "杰拉代",
		TitleCode: "b_jerada",
	}
}
