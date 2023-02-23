package tyrone

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 马赫拉MagheraBarony struct {
	feud.BaseBarony
}

var BMaghera马赫拉 feud.Barony = &马赫拉MagheraBarony{}

func init() {
	f := BMaghera马赫拉.(*马赫拉MagheraBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "maghera",
		TitleName: "马赫拉",
		TitleCode: "b_maghera",
	}
}
