package porto

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿尔库什迪瓦尔德维什ArcosdevaldevezBarony struct {
	feud.BaseBarony
}

var BArcosdevaldevez阿尔库什迪瓦尔德维什 feud.Barony = &阿尔库什迪瓦尔德维什ArcosdevaldevezBarony{}

func init() {
    f := BArcosdevaldevez阿尔库什迪瓦尔德维什.(*阿尔库什迪瓦尔德维什ArcosdevaldevezBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "arcosdevaldevez",
		TitleName: "阿尔库什迪瓦尔德维什",
		TitleCode: "b_arcosdevaldevez",
	}
}
