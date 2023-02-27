package kondana

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 军荼那KondanaBarony struct {
	feud.BaseBarony
}

var BKondana军荼那 feud.Barony = &军荼那KondanaBarony{}

func init() {
    f := BKondana军荼那.(*军荼那KondanaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kondana",
		TitleName: "军荼那",
		TitleCode: "b_kondana",
	}
}
