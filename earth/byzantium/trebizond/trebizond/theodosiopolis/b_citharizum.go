package theodosiopolis

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 基萨里宗CitharizumBarony struct {
	feud.BaseBarony
}

var BCitharizum基萨里宗 feud.Barony = &基萨里宗CitharizumBarony{}

func init() {
    f := BCitharizum基萨里宗.(*基萨里宗CitharizumBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "citharizum",
		TitleName: "基萨里宗",
		TitleCode: "b_citharizum",
	}
}
