package sarkel

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 白塔BelayavezhaBarony struct {
	feud.BaseBarony
}

var BBelayavezha白塔 feud.Barony = &白塔BelayavezhaBarony{}

func init() {
    f := BBelayavezha白塔.(*白塔BelayavezhaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "belayavezha",
		TitleName: "白塔",
		TitleCode: "b_belayavezha",
	}
}
