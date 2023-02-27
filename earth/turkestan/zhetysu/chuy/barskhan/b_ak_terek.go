package barskhan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿克铁列克Ak_terekBarony struct {
	feud.BaseBarony
}

var BAk_terek阿克铁列克 feud.Barony = &阿克铁列克Ak_terekBarony{}

func init() {
    f := BAk_terek阿克铁列克.(*阿克铁列克Ak_terekBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ak_terek",
		TitleName: "阿克铁列克",
		TitleCode: "b_ak_terek",
	}
}
