package araouane

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 廷扎乌登TinzaoutenBarony struct {
	feud.BaseBarony
}

var BTinzaouten廷扎乌登 feud.Barony = &廷扎乌登TinzaoutenBarony{}

func init() {
    f := BTinzaouten廷扎乌登.(*廷扎乌登TinzaoutenBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "tinzaouten",
		TitleName: "廷扎乌登",
		TitleCode: "b_tinzaouten",
	}
}
