package anti_atlas

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿克穆德AkermoudBarony struct {
	feud.BaseBarony
}

var BAkermoud阿克穆德 feud.Barony = &阿克穆德AkermoudBarony{}

func init() {
    f := BAkermoud阿克穆德.(*阿克穆德AkermoudBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "akermoud",
		TitleName: "阿克穆德",
		TitleCode: "b_akermoud",
	}
}
