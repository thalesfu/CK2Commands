package tom

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 热尔诺沃ZhernovoBarony struct {
	feud.BaseBarony
}

var BZhernovo热尔诺沃 feud.Barony = &热尔诺沃ZhernovoBarony{}

func init() {
	f := BZhernovo热尔诺沃.(*热尔诺沃ZhernovoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "zhernovo",
		TitleName: "热尔诺沃",
		TitleCode: "b_zhernovo",
	}
}
