package candhoba

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 求那GunaBarony struct {
	feud.BaseBarony
}

var BGuna求那 feud.Barony = &求那GunaBarony{}

func init() {
    f := BGuna求那.(*求那GunaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "guna",
		TitleName: "求那",
		TitleCode: "b_guna",
	}
}
