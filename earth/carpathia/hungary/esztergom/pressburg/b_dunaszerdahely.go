package pressburg

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 杜瑙塞尔道海伊DunaszerdahelyBarony struct {
	feud.BaseBarony
}

var BDunaszerdahely杜瑙塞尔道海伊 feud.Barony = &杜瑙塞尔道海伊DunaszerdahelyBarony{}

func init() {
    f := BDunaszerdahely杜瑙塞尔道海伊.(*杜瑙塞尔道海伊DunaszerdahelyBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "dunaszerdahely",
		TitleName: "杜瑙塞尔道海伊",
		TitleCode: "b_dunaszerdahely",
	}
}
