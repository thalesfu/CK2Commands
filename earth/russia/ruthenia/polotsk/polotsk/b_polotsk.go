package polotsk

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 波洛茨克PolotskBarony struct {
	feud.BaseBarony
}

var BPolotsk波洛茨克 feud.Barony = &波洛茨克PolotskBarony{}

func init() {
	f := BPolotsk波洛茨克.(*波洛茨克PolotskBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "polotsk",
		TitleName: "波洛茨克",
		TitleCode: "b_polotsk",
	}
}
