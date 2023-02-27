package vel

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 杜托沃DutovoBarony struct {
	feud.BaseBarony
}

var BDutovo杜托沃 feud.Barony = &杜托沃DutovoBarony{}

func init() {
    f := BDutovo杜托沃.(*杜托沃DutovoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "dutovo",
		TitleName: "杜托沃",
		TitleCode: "b_dutovo",
	}
}
