package bihar

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 瑙吉瓦劳德NagyvaradBarony struct {
	feud.BaseBarony
}

var BNagyvarad瑙吉瓦劳德 feud.Barony = &瑙吉瓦劳德NagyvaradBarony{}

func init() {
	f := BNagyvarad瑙吉瓦劳德.(*瑙吉瓦劳德NagyvaradBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "nagyvarad",
		TitleName: "瑙吉瓦劳德",
		TitleCode: "b_nagyvarad",
	}
}
