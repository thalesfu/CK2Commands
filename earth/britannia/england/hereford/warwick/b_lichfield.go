package warwick

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 利奇菲尔德LichfieldBarony struct {
	feud.BaseBarony
}

var BLichfield利奇菲尔德 feud.Barony = &利奇菲尔德LichfieldBarony{}

func init() {
    f := BLichfield利奇菲尔德.(*利奇菲尔德LichfieldBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "lichfield",
		TitleName: "利奇菲尔德",
		TitleCode: "b_lichfield",
	}
}
