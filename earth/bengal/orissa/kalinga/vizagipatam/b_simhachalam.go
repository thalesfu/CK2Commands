package vizagipatam

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 僧诃遮蓝SimhachalamBarony struct {
	feud.BaseBarony
}

var BSimhachalam僧诃遮蓝 feud.Barony = &僧诃遮蓝SimhachalamBarony{}

func init() {
	f := BSimhachalam僧诃遮蓝.(*僧诃遮蓝SimhachalamBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "simhachalam",
		TitleName: "僧诃遮蓝",
		TitleCode: "b_simhachalam",
	}
}
