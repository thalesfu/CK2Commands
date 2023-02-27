package luxembourg

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 讷沙托NeufchateauBarony struct {
	feud.BaseBarony
}

var BNeufchateau讷沙托 feud.Barony = &讷沙托NeufchateauBarony{}

func init() {
    f := BNeufchateau讷沙托.(*讷沙托NeufchateauBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "neufchateau",
		TitleName: "讷沙托",
		TitleCode: "b_neufchateau",
	}
}
