package astibus

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿斯提博斯AstibusBarony struct {
	feud.BaseBarony
}

var BAstibus阿斯提博斯 feud.Barony = &阿斯提博斯AstibusBarony{}

func init() {
	f := BAstibus阿斯提博斯.(*阿斯提博斯AstibusBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "astibus",
		TitleName: "阿斯提博斯",
		TitleCode: "b_astibus",
	}
}
