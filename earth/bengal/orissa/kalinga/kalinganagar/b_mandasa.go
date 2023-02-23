package kalinganagar

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 曼陀娑MandasaBarony struct {
	feud.BaseBarony
}

var BMandasa曼陀娑 feud.Barony = &曼陀娑MandasaBarony{}

func init() {
	f := BMandasa曼陀娑.(*曼陀娑MandasaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "mandasa",
		TitleName: "曼陀娑",
		TitleCode: "b_mandasa",
	}
}
