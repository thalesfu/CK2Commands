package quattara

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 济瓦巴哈里ZiwaelbahariBarony struct {
	feud.BaseBarony
}

var BZiwaelbahari济瓦巴哈里 feud.Barony = &济瓦巴哈里ZiwaelbahariBarony{}

func init() {
    f := BZiwaelbahari济瓦巴哈里.(*济瓦巴哈里ZiwaelbahariBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ziwaelbahari",
		TitleName: "济瓦巴哈里",
		TitleCode: "b_ziwaelbahari",
	}
}
