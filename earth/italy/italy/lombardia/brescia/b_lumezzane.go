package brescia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 卢梅扎内LumezzaneBarony struct {
	feud.BaseBarony
}

var BLumezzane卢梅扎内 feud.Barony = &卢梅扎内LumezzaneBarony{}

func init() {
    f := BLumezzane卢梅扎内.(*卢梅扎内LumezzaneBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "lumezzane",
		TitleName: "卢梅扎内",
		TitleCode: "b_lumezzane",
	}
}
