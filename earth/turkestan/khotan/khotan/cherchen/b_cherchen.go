package cherchen

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 车尔臣CherchenBarony struct {
	feud.BaseBarony
}

var BCherchen车尔臣 feud.Barony = &车尔臣CherchenBarony{}

func init() {
    f := BCherchen车尔臣.(*车尔臣CherchenBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "cherchen",
		TitleName: "车尔臣",
		TitleCode: "b_cherchen",
	}
}
