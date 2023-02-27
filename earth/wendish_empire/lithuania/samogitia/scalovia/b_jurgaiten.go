package scalovia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 尤尔盖滕JurgaitenBarony struct {
	feud.BaseBarony
}

var BJurgaiten尤尔盖滕 feud.Barony = &尤尔盖滕JurgaitenBarony{}

func init() {
    f := BJurgaiten尤尔盖滕.(*尤尔盖滕JurgaitenBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "jurgaiten",
		TitleName: "尤尔盖滕",
		TitleCode: "b_jurgaiten",
	}
}
