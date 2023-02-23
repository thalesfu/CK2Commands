package sikkim

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 曼甘ManganBarony struct {
	feud.BaseBarony
}

var BMangan曼甘 feud.Barony = &曼甘ManganBarony{}

func init() {
	f := BMangan曼甘.(*曼甘ManganBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "mangan",
		TitleName: "曼甘",
		TitleCode: "b_mangan",
	}
}
