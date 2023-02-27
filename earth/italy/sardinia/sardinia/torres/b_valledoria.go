package torres

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 瓦莱多里亚ValledoriaBarony struct {
	feud.BaseBarony
}

var BValledoria瓦莱多里亚 feud.Barony = &瓦莱多里亚ValledoriaBarony{}

func init() {
    f := BValledoria瓦莱多里亚.(*瓦莱多里亚ValledoriaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "valledoria",
		TitleName: "瓦莱多里亚",
		TitleCode: "b_valledoria",
	}
}
