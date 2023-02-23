package famagusta

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 凯里尼亚KyreniaBarony struct {
	feud.BaseBarony
}

var BKyrenia凯里尼亚 feud.Barony = &凯里尼亚KyreniaBarony{}

func init() {
	f := BKyrenia凯里尼亚.(*凯里尼亚KyreniaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kyrenia",
		TitleName: "凯里尼亚",
		TitleCode: "b_kyrenia",
	}
}
