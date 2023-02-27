package zhongba

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 亚热YagraBarony struct {
	feud.BaseBarony
}

var BYagra亚热 feud.Barony = &亚热YagraBarony{}

func init() {
    f := BYagra亚热.(*亚热YagraBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "yagra",
		TitleName: "亚热",
		TitleCode: "b_yagra",
	}
}
