package khaylam

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 纳吉姆NajmBarony struct {
	feud.BaseBarony
}

var BNajm纳吉姆 feud.Barony = &纳吉姆NajmBarony{}

func init() {
    f := BNajm纳吉姆.(*纳吉姆NajmBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "najm",
		TitleName: "纳吉姆",
		TitleCode: "b_najm",
	}
}
