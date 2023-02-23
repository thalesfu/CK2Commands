package nyima

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 塘来TanglaiBarony struct {
	feud.BaseBarony
}

var BTanglai塘来 feud.Barony = &塘来TanglaiBarony{}

func init() {
	f := BTanglai塘来.(*塘来TanglaiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "tanglai",
		TitleName: "塘来",
		TitleCode: "b_tanglai",
	}
}
