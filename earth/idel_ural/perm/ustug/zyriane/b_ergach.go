package zyriane

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 叶尔加奇ErgachBarony struct {
	feud.BaseBarony
}

var BErgach叶尔加奇 feud.Barony = &叶尔加奇ErgachBarony{}

func init() {
    f := BErgach叶尔加奇.(*叶尔加奇ErgachBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ergach",
		TitleName: "叶尔加奇",
		TitleCode: "b_ergach",
	}
}
