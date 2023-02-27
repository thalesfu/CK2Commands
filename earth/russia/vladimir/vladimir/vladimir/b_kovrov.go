package vladimir

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 科夫罗夫KovrovBarony struct {
	feud.BaseBarony
}

var BKovrov科夫罗夫 feud.Barony = &科夫罗夫KovrovBarony{}

func init() {
    f := BKovrov科夫罗夫.(*科夫罗夫KovrovBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kovrov",
		TitleName: "科夫罗夫",
		TitleCode: "b_kovrov",
	}
}
