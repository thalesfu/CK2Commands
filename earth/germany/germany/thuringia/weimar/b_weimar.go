package weimar

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 魏玛WeimarBarony struct {
	feud.BaseBarony
}

var BWeimar魏玛 feud.Barony = &魏玛WeimarBarony{}

func init() {
    f := BWeimar魏玛.(*魏玛WeimarBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "weimar",
		TitleName: "魏玛",
		TitleCode: "b_weimar",
	}
}
