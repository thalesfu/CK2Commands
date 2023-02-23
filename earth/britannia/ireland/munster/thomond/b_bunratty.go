package thomond

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 班拉蒂BunrattyBarony struct {
	feud.BaseBarony
}

var BBunratty班拉蒂 feud.Barony = &班拉蒂BunrattyBarony{}

func init() {
	f := BBunratty班拉蒂.(*班拉蒂BunrattyBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "bunratty",
		TitleName: "班拉蒂",
		TitleCode: "b_bunratty",
	}
}
