package torres

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 奥斯基里OschiriBarony struct {
	feud.BaseBarony
}

var BOschiri奥斯基里 feud.Barony = &奥斯基里OschiriBarony{}

func init() {
    f := BOschiri奥斯基里.(*奥斯基里OschiriBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "oschiri",
		TitleName: "奥斯基里",
		TitleCode: "b_oschiri",
	}
}
