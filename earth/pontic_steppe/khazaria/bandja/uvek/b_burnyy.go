package uvek

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 布尔内BurnyyBarony struct {
	feud.BaseBarony
}

var BBurnyy布尔内 feud.Barony = &布尔内BurnyyBarony{}

func init() {
    f := BBurnyy布尔内.(*布尔内BurnyyBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "burnyy",
		TitleName: "布尔内",
		TitleCode: "b_burnyy",
	}
}
