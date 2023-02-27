package ludrava

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 颇罗伐地迦PhalavardhikaBarony struct {
	feud.BaseBarony
}

var BPhalavardhika颇罗伐地迦 feud.Barony = &颇罗伐地迦PhalavardhikaBarony{}

func init() {
    f := BPhalavardhika颇罗伐地迦.(*颇罗伐地迦PhalavardhikaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "phalavardhika",
		TitleName: "颇罗伐地迦",
		TitleCode: "b_phalavardhika",
	}
}
