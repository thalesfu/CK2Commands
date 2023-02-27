package kafirkot

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 娑伽罗SagraBarony struct {
	feud.BaseBarony
}

var BSagra娑伽罗 feud.Barony = &娑伽罗SagraBarony{}

func init() {
    f := BSagra娑伽罗.(*娑伽罗SagraBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "sagra",
		TitleName: "娑伽罗",
		TitleCode: "b_sagra",
	}
}
