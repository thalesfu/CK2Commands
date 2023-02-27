package candhoba

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 帝卢僧耆TilosingiBarony struct {
	feud.BaseBarony
}

var BTilosingi帝卢僧耆 feud.Barony = &帝卢僧耆TilosingiBarony{}

func init() {
    f := BTilosingi帝卢僧耆.(*帝卢僧耆TilosingiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "tilosingi",
		TitleName: "帝卢僧耆",
		TitleCode: "b_tilosingi",
	}
}
