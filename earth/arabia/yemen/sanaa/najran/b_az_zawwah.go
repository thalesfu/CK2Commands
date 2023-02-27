package najran

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 扎乌瓦Az_zawwahBarony struct {
	feud.BaseBarony
}

var BAz_zawwah扎乌瓦 feud.Barony = &扎乌瓦Az_zawwahBarony{}

func init() {
    f := BAz_zawwah扎乌瓦.(*扎乌瓦Az_zawwahBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "az_zawwah",
		TitleName: "扎乌瓦",
		TitleCode: "b_az_zawwah",
	}
}
