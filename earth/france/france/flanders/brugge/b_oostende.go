package brugge

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 奥斯坦德OostendeBarony struct {
	feud.BaseBarony
}

var BOostende奥斯坦德 feud.Barony = &奥斯坦德OostendeBarony{}

func init() {
	f := BOostende奥斯坦德.(*奥斯坦德OostendeBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "oostende",
		TitleName: "奥斯坦德",
		TitleCode: "b_oostende",
	}
}
