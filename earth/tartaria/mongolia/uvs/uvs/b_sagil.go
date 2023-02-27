package uvs

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 萨吉勒SagilBarony struct {
	feud.BaseBarony
}

var BSagil萨吉勒 feud.Barony = &萨吉勒SagilBarony{}

func init() {
    f := BSagil萨吉勒.(*萨吉勒SagilBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "sagil",
		TitleName: "萨吉勒",
		TitleCode: "b_sagil",
	}
}
