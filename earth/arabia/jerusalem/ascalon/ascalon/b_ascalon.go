package ascalon

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 亚实基伦AscalonBarony struct {
	feud.BaseBarony
}

var BAscalon亚实基伦 feud.Barony = &亚实基伦AscalonBarony{}

func init() {
	f := BAscalon亚实基伦.(*亚实基伦AscalonBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ascalon",
		TitleName: "亚实基伦",
		TitleCode: "b_ascalon",
	}
}
