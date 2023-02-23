package vizagipatam

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 婆毗军荼BavikondaBarony struct {
	feud.BaseBarony
}

var BBavikonda婆毗军荼 feud.Barony = &婆毗军荼BavikondaBarony{}

func init() {
	f := BBavikonda婆毗军荼.(*婆毗军荼BavikondaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "bavikonda",
		TitleName: "婆毗军荼",
		TitleCode: "b_bavikonda",
	}
}
