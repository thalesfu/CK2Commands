package zagreb

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 采廷CetinBarony struct {
	feud.BaseBarony
}

var BCetin采廷 feud.Barony = &采廷CetinBarony{}

func init() {
    f := BCetin采廷.(*采廷CetinBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "cetin",
		TitleName: "采廷",
		TitleCode: "b_cetin",
	}
}
