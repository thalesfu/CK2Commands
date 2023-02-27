package jharkand

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 婆伽缚BhagvaBarony struct {
	feud.BaseBarony
}

var BBhagva婆伽缚 feud.Barony = &婆伽缚BhagvaBarony{}

func init() {
    f := BBhagva婆伽缚.(*婆伽缚BhagvaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "bhagva",
		TitleName: "婆伽缚",
		TitleCode: "b_bhagva",
	}
}
