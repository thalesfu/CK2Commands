package mohadavasaka

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 慕诃陀婆娑迦MohadavasakaBarony struct {
	feud.BaseBarony
}

var BMohadavasaka慕诃陀婆娑迦 feud.Barony = &慕诃陀婆娑迦MohadavasakaBarony{}

func init() {
    f := BMohadavasaka慕诃陀婆娑迦.(*慕诃陀婆娑迦MohadavasakaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "mohadavasaka",
		TitleName: "慕诃陀婆娑迦",
		TitleCode: "b_mohadavasaka",
	}
}
