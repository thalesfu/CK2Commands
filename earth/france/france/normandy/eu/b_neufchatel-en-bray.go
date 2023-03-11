package eu

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 讷沙泰勒昂布赖Neufchatel_en_brayBarony struct {
	feud.BaseBarony
}

var BNeufchatel_en_bray讷沙泰勒昂布赖 feud.Barony = &讷沙泰勒昂布赖Neufchatel_en_brayBarony{}

func init() {
    f := BNeufchatel_en_bray讷沙泰勒昂布赖.(*讷沙泰勒昂布赖Neufchatel_en_brayBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "neufchatel_en_bray",
		TitleName: "讷沙泰勒昂布赖",
		TitleCode: "b_neufchatel-en-bray",
	}
}
