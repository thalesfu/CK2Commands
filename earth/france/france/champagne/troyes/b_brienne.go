package troyes

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 布列讷BrienneBarony struct {
	feud.BaseBarony
}

var BBrienne布列讷 feud.Barony = &布列讷BrienneBarony{}

func init() {
	f := BBrienne布列讷.(*布列讷BrienneBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "brienne",
		TitleName: "布列讷",
		TitleCode: "b_brienne",
	}
}
