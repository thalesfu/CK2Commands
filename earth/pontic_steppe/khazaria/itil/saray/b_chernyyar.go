package saray

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 黑峡谷ChernyyarBarony struct {
	feud.BaseBarony
}

var BChernyyar黑峡谷 feud.Barony = &黑峡谷ChernyyarBarony{}

func init() {
    f := BChernyyar黑峡谷.(*黑峡谷ChernyyarBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "chernyyar",
		TitleName: "黑峡谷",
		TitleCode: "b_chernyyar",
	}
}
