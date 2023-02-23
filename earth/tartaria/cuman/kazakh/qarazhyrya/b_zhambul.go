package qarazhyrya

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 江布尔ZhambulBarony struct {
	feud.BaseBarony
}

var BZhambul江布尔 feud.Barony = &江布尔ZhambulBarony{}

func init() {
	f := BZhambul江布尔.(*江布尔ZhambulBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "zhambul",
		TitleName: "江布尔",
		TitleCode: "b_zhambul",
	}
}
