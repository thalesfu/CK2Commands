package qalqut

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 内迪耶卢补NediyiruppuBarony struct {
	feud.BaseBarony
}

var BNediyiruppu内迪耶卢补 feud.Barony = &内迪耶卢补NediyiruppuBarony{}

func init() {
    f := BNediyiruppu内迪耶卢补.(*内迪耶卢补NediyiruppuBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "nediyiruppu",
		TitleName: "内迪耶卢补",
		TitleCode: "b_nediyiruppu",
	}
}
