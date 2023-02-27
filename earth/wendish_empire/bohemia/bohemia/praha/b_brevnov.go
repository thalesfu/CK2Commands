package praha

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 布热夫诺夫BrevnovBarony struct {
	feud.BaseBarony
}

var BBrevnov布热夫诺夫 feud.Barony = &布热夫诺夫BrevnovBarony{}

func init() {
    f := BBrevnov布热夫诺夫.(*布热夫诺夫BrevnovBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "brevnov",
		TitleName: "布热夫诺夫",
		TitleCode: "b_brevnov",
	}
}
