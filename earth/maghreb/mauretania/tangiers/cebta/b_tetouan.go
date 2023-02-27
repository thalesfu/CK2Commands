package cebta

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 得土安TetouanBarony struct {
	feud.BaseBarony
}

var BTetouan得土安 feud.Barony = &得土安TetouanBarony{}

func init() {
    f := BTetouan得土安.(*得土安TetouanBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "tetouan",
		TitleName: "得土安",
		TitleCode: "b_tetouan",
	}
}
