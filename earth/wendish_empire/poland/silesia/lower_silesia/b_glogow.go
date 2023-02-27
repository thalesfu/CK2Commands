package lower_silesia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 格沃古夫GlogowBarony struct {
	feud.BaseBarony
}

var BGlogow格沃古夫 feud.Barony = &格沃古夫GlogowBarony{}

func init() {
    f := BGlogow格沃古夫.(*格沃古夫GlogowBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "glogow",
		TitleName: "格沃古夫",
		TitleCode: "b_glogow",
	}
}
