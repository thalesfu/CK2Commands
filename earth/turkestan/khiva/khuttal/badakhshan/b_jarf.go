package badakhshan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 贾尔夫JarfBarony struct {
	feud.BaseBarony
}

var BJarf贾尔夫 feud.Barony = &贾尔夫JarfBarony{}

func init() {
	f := BJarf贾尔夫.(*贾尔夫JarfBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "jarf",
		TitleName: "贾尔夫",
		TitleCode: "b_jarf",
	}
}
