package naissus

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 库马诺沃KumanovoBarony struct {
	feud.BaseBarony
}

var BKumanovo库马诺沃 feud.Barony = &库马诺沃KumanovoBarony{}

func init() {
	f := BKumanovo库马诺沃.(*库马诺沃KumanovoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kumanovo",
		TitleName: "库马诺沃",
		TitleCode: "b_kumanovo",
	}
}
