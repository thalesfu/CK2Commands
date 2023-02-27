package sortavala

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 拉彭兰塔LappeenrantaBarony struct {
	feud.BaseBarony
}

var BLappeenranta拉彭兰塔 feud.Barony = &拉彭兰塔LappeenrantaBarony{}

func init() {
    f := BLappeenranta拉彭兰塔.(*拉彭兰塔LappeenrantaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "lappeenranta",
		TitleName: "拉彭兰塔",
		TitleCode: "b_lappeenranta",
	}
}
