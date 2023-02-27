package almaty

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿克札尔Akzhar_almatyBarony struct {
	feud.BaseBarony
}

var BAkzhar_almaty阿克札尔 feud.Barony = &阿克札尔Akzhar_almatyBarony{}

func init() {
    f := BAkzhar_almaty阿克札尔.(*阿克札尔Akzhar_almatyBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "akzhar_almaty",
		TitleName: "阿克札尔",
		TitleCode: "b_akzhar_almaty",
	}
}
