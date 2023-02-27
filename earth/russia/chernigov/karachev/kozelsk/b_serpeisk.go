package kozelsk

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 谢尔佩伊斯克SerpeiskBarony struct {
	feud.BaseBarony
}

var BSerpeisk谢尔佩伊斯克 feud.Barony = &谢尔佩伊斯克SerpeiskBarony{}

func init() {
    f := BSerpeisk谢尔佩伊斯克.(*谢尔佩伊斯克SerpeiskBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "serpeisk",
		TitleName: "谢尔佩伊斯克",
		TitleCode: "b_serpeisk",
	}
}
