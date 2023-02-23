package hedmark

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 哈马尔HamarBarony struct {
	feud.BaseBarony
}

var BHamar哈马尔 feud.Barony = &哈马尔HamarBarony{}

func init() {
	f := BHamar哈马尔.(*哈马尔HamarBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "hamar",
		TitleName: "哈马尔",
		TitleCode: "b_hamar",
	}
}
