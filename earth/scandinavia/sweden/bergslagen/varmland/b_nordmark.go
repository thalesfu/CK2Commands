package varmland

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 努德马克NordmarkBarony struct {
	feud.BaseBarony
}

var BNordmark努德马克 feud.Barony = &努德马克NordmarkBarony{}

func init() {
	f := BNordmark努德马克.(*努德马克NordmarkBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "nordmark",
		TitleName: "努德马克",
		TitleCode: "b_nordmark",
	}
}
