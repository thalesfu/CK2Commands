package shrewsbury

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 文洛克WenlockBarony struct {
	feud.BaseBarony
}

var BWenlock文洛克 feud.Barony = &文洛克WenlockBarony{}

func init() {
    f := BWenlock文洛克.(*文洛克WenlockBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "wenlock",
		TitleName: "文洛克",
		TitleCode: "b_wenlock",
	}
}
