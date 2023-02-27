package kanin

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 奥泽罗OzeroBarony struct {
	feud.BaseBarony
}

var BOzero奥泽罗 feud.Barony = &奥泽罗OzeroBarony{}

func init() {
    f := BOzero奥泽罗.(*奥泽罗OzeroBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ozero",
		TitleName: "奥泽罗",
		TitleCode: "b_ozero",
	}
}
