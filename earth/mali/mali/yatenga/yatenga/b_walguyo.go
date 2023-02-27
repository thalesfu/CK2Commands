package yatenga

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 瓦勒古约WalguyoBarony struct {
	feud.BaseBarony
}

var BWalguyo瓦勒古约 feud.Barony = &瓦勒古约WalguyoBarony{}

func init() {
    f := BWalguyo瓦勒古约.(*瓦勒古约WalguyoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "walguyo",
		TitleName: "瓦勒古约",
		TitleCode: "b_walguyo",
	}
}
