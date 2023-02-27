package penugonda

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 都迦罗军吒TogarakuntaBarony struct {
	feud.BaseBarony
}

var BTogarakunta都迦罗军吒 feud.Barony = &都迦罗军吒TogarakuntaBarony{}

func init() {
    f := BTogarakunta都迦罗军吒.(*都迦罗军吒TogarakuntaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "togarakunta",
		TitleName: "都迦罗军吒",
		TitleCode: "b_togarakunta",
	}
}
