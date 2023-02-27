package khijjingakota

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 揭吒伽罗摩GhatagaonBarony struct {
	feud.BaseBarony
}

var BGhatagaon揭吒伽罗摩 feud.Barony = &揭吒伽罗摩GhatagaonBarony{}

func init() {
    f := BGhatagaon揭吒伽罗摩.(*揭吒伽罗摩GhatagaonBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ghatagaon",
		TitleName: "揭吒伽罗摩",
		TitleCode: "b_ghatagaon",
	}
}
