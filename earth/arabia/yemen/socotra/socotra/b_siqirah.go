package socotra

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 锡吉拉SiqirahBarony struct {
	feud.BaseBarony
}

var BSiqirah锡吉拉 feud.Barony = &锡吉拉SiqirahBarony{}

func init() {
	f := BSiqirah锡吉拉.(*锡吉拉SiqirahBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "siqirah",
		TitleName: "锡吉拉",
		TitleCode: "b_siqirah",
	}
}
