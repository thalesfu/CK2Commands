package juyan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type JuyanCounty interface {
    feud.County
    BChuyen_hai居延海() 	feud.Barony
    BGaxun_nur嘎顺淖尔() 	feud.Barony
    BJuyan居延() 	feud.Barony
    BMurengaole木仁高勒() 	feud.Barony
    BShivee_khuren西伯库伦() 	feud.Barony
    BSogo_nuur索果淖尔() 	feud.Barony
    BSubozhuoer苏泊淖尔() 	feud.Barony
}

type 居延JuyanCounty struct {
	feud.BaseCounty
	居延海Chuyen_hai 	feud.Barony
	嘎顺淖尔Gaxun_nur 	feud.Barony
	居延Juyan 	feud.Barony
	木仁高勒Murengaole 	feud.Barony
	西伯库伦Shivee_khuren 	feud.Barony
	索果淖尔Sogo_nuur 	feud.Barony
	苏泊淖尔Subozhuoer 	feud.Barony
}

func (c *居延JuyanCounty) BChuyen_hai居延海() feud.Barony {
	return c.居延海Chuyen_hai
}
    
func (c *居延JuyanCounty) BGaxun_nur嘎顺淖尔() feud.Barony {
	return c.嘎顺淖尔Gaxun_nur
}
    
func (c *居延JuyanCounty) BJuyan居延() feud.Barony {
	return c.居延Juyan
}
    
func (c *居延JuyanCounty) BMurengaole木仁高勒() feud.Barony {
	return c.木仁高勒Murengaole
}
    
func (c *居延JuyanCounty) BShivee_khuren西伯库伦() feud.Barony {
	return c.西伯库伦Shivee_khuren
}
    
func (c *居延JuyanCounty) BSogo_nuur索果淖尔() feud.Barony {
	return c.索果淖尔Sogo_nuur
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

	f.居延海Chuyen_hai = BChuyen_hai居延海
	f.居延海Chuyen_hai.SetParent(f)
	
	f.嘎顺淖尔Gaxun_nur = BGaxun_nur嘎顺淖尔
	f.嘎顺淖尔Gaxun_nur.SetParent(f)
	
	f.居延Juyan = BJuyan居延
	f.居延Juyan.SetParent(f)
	
	f.木仁高勒Murengaole = BMurengaole木仁高勒
	f.木仁高勒Murengaole.SetParent(f)
	
	f.西伯库伦Shivee_khuren = BShivee_khuren西伯库伦
	f.西伯库伦Shivee_khuren.SetParent(f)
	
	f.索果淖尔Sogo_nuur = BSogo_nuur索果淖尔
	f.索果淖尔Sogo_nuur.SetParent(f)
	
	f.苏泊淖尔Subozhuoer = BSubozhuoer苏泊淖尔
	f.苏泊淖尔Subozhuoer.SetParent(f)
	
}
