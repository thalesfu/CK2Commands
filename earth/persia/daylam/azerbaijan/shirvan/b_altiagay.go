package shirvan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿勒提干AltiagayBarony struct {
	feud.BaseBarony
}

var BAltiagay阿勒提干 feud.Barony = &阿勒提干AltiagayBarony{}

func init() {
    f := BAltiagay阿勒提干.(*阿勒提干AltiagayBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "altiagay",
		TitleName: "阿勒提干",
		TitleCode: "b_altiagay",
	}
}
