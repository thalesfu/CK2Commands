package khijjingakota

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 吉里安纳加尔KirianagarBarony struct {
	feud.BaseBarony
}

var BKirianagar吉里安纳加尔 feud.Barony = &吉里安纳加尔KirianagarBarony{}

func init() {
    f := BKirianagar吉里安纳加尔.(*吉里安纳加尔KirianagarBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kirianagar",
		TitleName: "吉里安纳加尔",
		TitleCode: "b_kirianagar",
	}
}
