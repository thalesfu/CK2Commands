package ostergotland

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 乌特什贝里UttersbergBarony struct {
	feud.BaseBarony
}

var BUttersberg乌特什贝里 feud.Barony = &乌特什贝里UttersbergBarony{}

func init() {
    f := BUttersberg乌特什贝里.(*乌特什贝里UttersbergBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "uttersberg",
		TitleName: "乌特什贝里",
		TitleCode: "b_uttersberg",
	}
}
