package al_amarah

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 瓦斯特WasitBarony struct {
	feud.BaseBarony
}

var BWasit瓦斯特 feud.Barony = &瓦斯特WasitBarony{}

func init() {
    f := BWasit瓦斯特.(*瓦斯特WasitBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "wasit",
		TitleName: "瓦斯特",
		TitleCode: "b_wasit",
	}
}
