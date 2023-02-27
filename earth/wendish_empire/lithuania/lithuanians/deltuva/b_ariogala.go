package deltuva

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿廖加拉AriogalaBarony struct {
	feud.BaseBarony
}

var BAriogala阿廖加拉 feud.Barony = &阿廖加拉AriogalaBarony{}

func init() {
    f := BAriogala阿廖加拉.(*阿廖加拉AriogalaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ariogala",
		TitleName: "阿廖加拉",
		TitleCode: "b_ariogala",
	}
}
