package hamburg

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 布克斯特胡德BuxtehudeBarony struct {
	feud.BaseBarony
}

var BBuxtehude布克斯特胡德 feud.Barony = &布克斯特胡德BuxtehudeBarony{}

func init() {
    f := BBuxtehude布克斯特胡德.(*布克斯特胡德BuxtehudeBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "buxtehude",
		TitleName: "布克斯特胡德",
		TitleCode: "b_buxtehude",
	}
}
