package bytow

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 斯沃沃夫SwolowBarony struct {
	feud.BaseBarony
}

var BSwolow斯沃沃夫 feud.Barony = &斯沃沃夫SwolowBarony{}

func init() {
    f := BSwolow斯沃沃夫.(*斯沃沃夫SwolowBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "swolow",
		TitleName: "斯沃沃夫",
		TitleCode: "b_swolow",
	}
}
