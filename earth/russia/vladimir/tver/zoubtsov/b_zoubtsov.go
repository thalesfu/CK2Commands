package zoubtsov

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 祖布佐夫ZoubtsovBarony struct {
	feud.BaseBarony
}

var BZoubtsov祖布佐夫 feud.Barony = &祖布佐夫ZoubtsovBarony{}

func init() {
	f := BZoubtsov祖布佐夫.(*祖布佐夫ZoubtsovBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "zoubtsov",
		TitleName: "祖布佐夫",
		TitleCode: "b_zoubtsov",
	}
}
