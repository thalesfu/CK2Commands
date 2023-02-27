package valladolid

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 佩尼亚菲耶尔PenafielBarony struct {
	feud.BaseBarony
}

var BPenafiel佩尼亚菲耶尔 feud.Barony = &佩尼亚菲耶尔PenafielBarony{}

func init() {
    f := BPenafiel佩尼亚菲耶尔.(*佩尼亚菲耶尔PenafielBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "penafiel",
		TitleName: "佩尼亚菲耶尔",
		TitleCode: "b_penafiel",
	}
}
