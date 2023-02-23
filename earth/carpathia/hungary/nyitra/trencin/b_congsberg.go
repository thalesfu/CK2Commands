package trencin

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 孔斯堡CongsbergBarony struct {
	feud.BaseBarony
}

var BCongsberg孔斯堡 feud.Barony = &孔斯堡CongsbergBarony{}

func init() {
	f := BCongsberg孔斯堡.(*孔斯堡CongsbergBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "congsberg",
		TitleName: "孔斯堡",
		TitleCode: "b_congsberg",
	}
}
