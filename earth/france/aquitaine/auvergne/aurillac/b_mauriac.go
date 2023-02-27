package aurillac

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 莫里亚克MauriacBarony struct {
	feud.BaseBarony
}

var BMauriac莫里亚克 feud.Barony = &莫里亚克MauriacBarony{}

func init() {
    f := BMauriac莫里亚克.(*莫里亚克MauriacBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "mauriac",
		TitleName: "莫里亚克",
		TitleCode: "b_mauriac",
	}
}
