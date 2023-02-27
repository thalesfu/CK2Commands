package mughalzhar

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 坎德阿加什KandyagashBarony struct {
	feud.BaseBarony
}

var BKandyagash坎德阿加什 feud.Barony = &坎德阿加什KandyagashBarony{}

func init() {
    f := BKandyagash坎德阿加什.(*坎德阿加什KandyagashBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kandyagash",
		TitleName: "坎德阿加什",
		TitleCode: "b_kandyagash",
	}
}
