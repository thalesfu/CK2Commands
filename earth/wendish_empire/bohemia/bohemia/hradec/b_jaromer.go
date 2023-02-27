package hradec

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 亚罗梅日JaromerBarony struct {
	feud.BaseBarony
}

var BJaromer亚罗梅日 feud.Barony = &亚罗梅日JaromerBarony{}

func init() {
    f := BJaromer亚罗梅日.(*亚罗梅日JaromerBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "jaromer",
		TitleName: "亚罗梅日",
		TitleCode: "b_jaromer",
	}
}
