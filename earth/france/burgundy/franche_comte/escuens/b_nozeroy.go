package escuens

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 诺兹鲁瓦NozeroyBarony struct {
	feud.BaseBarony
}

var BNozeroy诺兹鲁瓦 feud.Barony = &诺兹鲁瓦NozeroyBarony{}

func init() {
    f := BNozeroy诺兹鲁瓦.(*诺兹鲁瓦NozeroyBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "nozeroy",
		TitleName: "诺兹鲁瓦",
		TitleCode: "b_nozeroy",
	}
}
