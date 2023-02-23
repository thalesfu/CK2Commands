package schwaben

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 霍恩贝格HohenbergBarony struct {
	feud.BaseBarony
}

var BHohenberg霍恩贝格 feud.Barony = &霍恩贝格HohenbergBarony{}

func init() {
	f := BHohenberg霍恩贝格.(*霍恩贝格HohenbergBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "hohenberg",
		TitleName: "霍恩贝格",
		TitleCode: "b_hohenberg",
	}
}
