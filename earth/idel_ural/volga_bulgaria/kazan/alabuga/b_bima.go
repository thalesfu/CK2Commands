package alabuga

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 比马BimaBarony struct {
	feud.BaseBarony
}

var BBima比马 feud.Barony = &比马BimaBarony{}

func init() {
    f := BBima比马.(*比马BimaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "bima",
		TitleName: "比马",
		TitleCode: "b_bima",
	}
}
