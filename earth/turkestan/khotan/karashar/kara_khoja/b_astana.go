package kara_khoja

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿斯塔那AstanaBarony struct {
	feud.BaseBarony
}

var BAstana阿斯塔那 feud.Barony = &阿斯塔那AstanaBarony{}

func init() {
    f := BAstana阿斯塔那.(*阿斯塔那AstanaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "astana",
		TitleName: "阿斯塔那",
		TitleCode: "b_astana",
	}
}
