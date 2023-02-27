package dyamare

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 伽亚Gaya_songhayBarony struct {
	feud.BaseBarony
}

var BGaya_songhay伽亚 feud.Barony = &伽亚Gaya_songhayBarony{}

func init() {
    f := BGaya_songhay伽亚.(*伽亚Gaya_songhayBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "gaya_songhay",
		TitleName: "伽亚",
		TitleCode: "b_gaya_songhay",
	}
}
