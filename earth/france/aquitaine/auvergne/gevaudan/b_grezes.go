package gevaudan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 格雷泽GrezesBarony struct {
	feud.BaseBarony
}

var BGrezes格雷泽 feud.Barony = &格雷泽GrezesBarony{}

func init() {
    f := BGrezes格雷泽.(*格雷泽GrezesBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "grezes",
		TitleName: "格雷泽",
		TitleCode: "b_grezes",
	}
}
