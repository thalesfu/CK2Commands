package escuens

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 伯兰BelinBarony struct {
	feud.BaseBarony
}

var BBelin伯兰 feud.Barony = &伯兰BelinBarony{}

func init() {
    f := BBelin伯兰.(*伯兰BelinBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "belin",
		TitleName: "伯兰",
		TitleCode: "b_belin",
	}
}
