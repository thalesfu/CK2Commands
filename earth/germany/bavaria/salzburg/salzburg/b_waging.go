package salzburg

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 瓦京WagingBarony struct {
	feud.BaseBarony
}

var BWaging瓦京 feud.Barony = &瓦京WagingBarony{}

func init() {
	f := BWaging瓦京.(*瓦京WagingBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "waging",
		TitleName: "瓦京",
		TitleCode: "b_waging",
	}
}
