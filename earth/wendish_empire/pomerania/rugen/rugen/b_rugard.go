package rugen

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 卢加德RugardBarony struct {
	feud.BaseBarony
}

var BRugard卢加德 feud.Barony = &卢加德RugardBarony{}

func init() {
    f := BRugard卢加德.(*卢加德RugardBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "rugard",
		TitleName: "卢加德",
		TitleCode: "b_rugard",
	}
}
