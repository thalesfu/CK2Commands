package leon

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 萨阿贡SahagunBarony struct {
	feud.BaseBarony
}

var BSahagun萨阿贡 feud.Barony = &萨阿贡SahagunBarony{}

func init() {
    f := BSahagun萨阿贡.(*萨阿贡SahagunBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "sahagun",
		TitleName: "萨阿贡",
		TitleCode: "b_sahagun",
	}
}
