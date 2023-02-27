package senoussi

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 贾哈尔拉JakharrahBarony struct {
	feud.BaseBarony
}

var BJakharrah贾哈尔拉 feud.Barony = &贾哈尔拉JakharrahBarony{}

func init() {
    f := BJakharrah贾哈尔拉.(*贾哈尔拉JakharrahBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "jakharrah",
		TitleName: "贾哈尔拉",
		TitleCode: "b_jakharrah",
	}
}
