package malta

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 博尔戈MarsascalaBarony struct {
	feud.BaseBarony
}

var BMarsascala博尔戈 feud.Barony = &博尔戈MarsascalaBarony{}

func init() {
    f := BMarsascala博尔戈.(*博尔戈MarsascalaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "marsascala",
		TitleName: "博尔戈",
		TitleCode: "b_marsascala",
	}
}
