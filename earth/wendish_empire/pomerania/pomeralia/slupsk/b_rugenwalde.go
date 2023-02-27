package slupsk

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 吕根瓦尔德RugenwaldeBarony struct {
	feud.BaseBarony
}

var BRugenwalde吕根瓦尔德 feud.Barony = &吕根瓦尔德RugenwaldeBarony{}

func init() {
    f := BRugenwalde吕根瓦尔德.(*吕根瓦尔德RugenwaldeBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "rugenwalde",
		TitleName: "吕根瓦尔德",
		TitleCode: "b_rugenwalde",
	}
}
