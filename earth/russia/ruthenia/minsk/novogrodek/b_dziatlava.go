package novogrodek

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 佳特洛沃DziatlavaBarony struct {
	feud.BaseBarony
}

var BDziatlava佳特洛沃 feud.Barony = &佳特洛沃DziatlavaBarony{}

func init() {
    f := BDziatlava佳特洛沃.(*佳特洛沃DziatlavaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "dziatlava",
		TitleName: "佳特洛沃",
		TitleCode: "b_dziatlava",
	}
}
