package juyan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type JuyanCounty interface {
	feud.County
	BJuyan居延() feud.Barony
	BMurengaole木仁高勒() feud.Barony
	BSubozhuoer苏泊淖尔() feud.Barony
}

type 居延JuyanCounty struct {
	feud.BaseCounty
	居延Juyan        feud.Barony
	木仁高勒Murengaole feud.Barony
	苏泊淖尔Subozhuoer feud.Barony
}

func (c *居延JuyanCounty) BJuyan居延() feud.Barony {
	return c.居延Juyan
}

func (c *居延JuyanCounty) BMurengaole木仁高勒() feud.Barony {
	return c.木仁高勒Murengaole
}

func (c *居延JuyanCounty) BSubozhuoer苏泊淖尔() feud.Barony {
	return c.苏泊淖尔Subozhuoer
}

var CJuyan居延 JuyanCounty = &居延JuyanCounty{}

func init() {
	f := CJuyan居延.(*居延JuyanCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1919",
		Title:     "juyan",
		TitleName: "居延",
		TitleCode: "c_juyan",
		Baronies:  map[string]feud.Barony{},
	}

	f.居延Juyan = BJuyan居延
	f.居延Juyan.SetParent(f)

	f.木仁高勒Murengaole = BMurengaole木仁高勒
	f.木仁高勒Murengaole.SetParent(f)

	f.苏泊淖尔Subozhuoer = BSubozhuoer苏泊淖尔
	f.苏泊淖尔Subozhuoer.SetParent(f)

}
