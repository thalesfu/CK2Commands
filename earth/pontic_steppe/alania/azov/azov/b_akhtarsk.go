package azov

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿赫塔尔斯克AkhtarskBarony struct {
	feud.BaseBarony
}

var BAkhtarsk阿赫塔尔斯克 feud.Barony = &阿赫塔尔斯克AkhtarskBarony{}

func init() {
    f := BAkhtarsk阿赫塔尔斯克.(*阿赫塔尔斯克AkhtarskBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "akhtarsk",
		TitleName: "阿赫塔尔斯克",
		TitleCode: "b_akhtarsk",
	}
}
