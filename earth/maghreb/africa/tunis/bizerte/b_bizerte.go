package bizerte

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 比塞大BizerteBarony struct {
	feud.BaseBarony
}

var BBizerte比塞大 feud.Barony = &比塞大BizerteBarony{}

func init() {
    f := BBizerte比塞大.(*比塞大BizerteBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "bizerte",
		TitleName: "比塞大",
		TitleCode: "b_bizerte",
	}
}
