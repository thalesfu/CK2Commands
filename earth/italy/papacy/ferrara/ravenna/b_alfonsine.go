package ravenna

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿方西内AlfonsineBarony struct {
	feud.BaseBarony
}

var BAlfonsine阿方西内 feud.Barony = &阿方西内AlfonsineBarony{}

func init() {
    f := BAlfonsine阿方西内.(*阿方西内AlfonsineBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "alfonsine",
		TitleName: "阿方西内",
		TitleCode: "b_alfonsine",
	}
}
