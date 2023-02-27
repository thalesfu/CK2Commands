package medog

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 甘登GandainBarony struct {
	feud.BaseBarony
}

var BGandain甘登 feud.Barony = &甘登GandainBarony{}

func init() {
    f := BGandain甘登.(*甘登GandainBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "gandain",
		TitleName: "甘登",
		TitleCode: "b_gandain",
	}
}
