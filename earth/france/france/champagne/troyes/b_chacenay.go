package troyes

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 沙瑟奈ChacenayBarony struct {
	feud.BaseBarony
}

var BChacenay沙瑟奈 feud.Barony = &沙瑟奈ChacenayBarony{}

func init() {
    f := BChacenay沙瑟奈.(*沙瑟奈ChacenayBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "chacenay",
		TitleName: "沙瑟奈",
		TitleCode: "b_chacenay",
	}
}
