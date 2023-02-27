package orleans

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 雅尔若JargeauBarony struct {
	feud.BaseBarony
}

var BJargeau雅尔若 feud.Barony = &雅尔若JargeauBarony{}

func init() {
    f := BJargeau雅尔若.(*雅尔若JargeauBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "jargeau",
		TitleName: "雅尔若",
		TitleCode: "b_jargeau",
	}
}
