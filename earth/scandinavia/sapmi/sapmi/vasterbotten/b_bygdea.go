package vasterbotten

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 比格迪奥BygdeaBarony struct {
	feud.BaseBarony
}

var BBygdea比格迪奥 feud.Barony = &比格迪奥BygdeaBarony{}

func init() {
    f := BBygdea比格迪奥.(*比格迪奥BygdeaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "bygdea",
		TitleName: "比格迪奥",
		TitleCode: "b_bygdea",
	}
}
