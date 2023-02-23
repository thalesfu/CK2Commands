package bath

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 布里斯托BristolBarony struct {
	feud.BaseBarony
}

var BBristol布里斯托 feud.Barony = &布里斯托BristolBarony{}

func init() {
	f := BBristol布里斯托.(*布里斯托BristolBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "bristol",
		TitleName: "布里斯托",
		TitleCode: "b_bristol",
	}
}
