package krakowskie

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 延杰尤夫JedrzejowBarony struct {
	feud.BaseBarony
}

var BJedrzejow延杰尤夫 feud.Barony = &延杰尤夫JedrzejowBarony{}

func init() {
    f := BJedrzejow延杰尤夫.(*延杰尤夫JedrzejowBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "jedrzejow",
		TitleName: "延杰尤夫",
		TitleCode: "b_jedrzejow",
	}
}
