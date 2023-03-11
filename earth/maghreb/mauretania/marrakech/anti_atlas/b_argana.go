package anti_atlas

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 艾尔加奈ArganaBarony struct {
	feud.BaseBarony
}

var BArgana艾尔加奈 feud.Barony = &艾尔加奈ArganaBarony{}

func init() {
    f := BArgana艾尔加奈.(*艾尔加奈ArganaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "argana",
		TitleName: "艾尔加奈",
		TitleCode: "b_argana",
	}
}
