package consenza

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 斯特龙戈利StrongoliBarony struct {
	feud.BaseBarony
}

var BStrongoli斯特龙戈利 feud.Barony = &斯特龙戈利StrongoliBarony{}

func init() {
	f := BStrongoli斯特龙戈利.(*斯特龙戈利StrongoliBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "strongoli",
		TitleName: "斯特龙戈利",
		TitleCode: "b_strongoli",
	}
}
