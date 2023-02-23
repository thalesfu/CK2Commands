package dunkheger

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 横岭HenglingBarony struct {
	feud.BaseBarony
}

var BHengling横岭 feud.Barony = &横岭HenglingBarony{}

func init() {
	f := BHengling横岭.(*横岭HenglingBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "hengling",
		TitleName: "横岭",
		TitleCode: "b_hengling",
	}
}
