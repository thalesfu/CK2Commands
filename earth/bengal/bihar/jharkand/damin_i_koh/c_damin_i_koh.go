package damin_i_koh

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type Damin_i_kohCounty interface {
    feud.County
    BBhadalthua婆荼罗土诃() 	feud.Barony
    BBhadua婆墯() 	feud.Barony
    BDhanaura陀奴罗() 	feud.Barony
    BDharau陀卢() 	feud.Barony
    BKannpur乾补() 	feud.Barony
    BShikarji尸佉罗时() 	feud.Barony
}

type 昙民Damin_i_kohCounty struct {
	feud.BaseCounty
	婆荼罗土诃Bhadalthua 	feud.Barony
	婆墯Bhadua 	feud.Barony
	陀奴罗Dhanaura 	feud.Barony
	陀卢Dharau 	feud.Barony
	乾补Kannpur 	feud.Barony
	尸佉罗时Shikarji 	feud.Barony
}

func (c *昙民Damin_i_kohCounty) BBhadalthua婆荼罗土诃() feud.Barony {
	return c.婆荼罗土诃Bhadalthua
}
    
func (c *昙民Damin_i_kohCounty) BBhadua婆墯() feud.Barony {
	return c.婆墯Bhadua
}
    
func (c *昙民Damin_i_kohCounty) BDhanaura陀奴罗() feud.Barony {
	return c.陀奴罗Dhanaura
}
    
func (c *昙民Damin_i_kohCounty) BDharau陀卢() feud.Barony {
	return c.陀卢Dharau
}
    
func (c *昙民Damin_i_kohCounty) BKannpur乾补() feud.Barony {
	return c.乾补Kannpur
}
    
func (c *昙民Damin_i_kohCounty) BShikarji尸佉罗时() feud.Barony {
	return c.尸佉罗时Shikarji
}
    
var CDamin_i_koh昙民 Damin_i_kohCounty = &昙民Damin_i_kohCounty{}

func init() {
	f := CDamin_i_koh昙民.(*昙民Damin_i_kohCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1234",
		Title:     "damin_i_koh",
		TitleName: "昙民",
		TitleCode: "c_damin_i_koh",
		Baronies:  map[string]feud.Barony{},
	}

	f.婆荼罗土诃Bhadalthua = BBhadalthua婆荼罗土诃
	f.婆荼罗土诃Bhadalthua.SetParent(f)
	
	f.婆墯Bhadua = BBhadua婆墯
	f.婆墯Bhadua.SetParent(f)
	
	f.陀奴罗Dhanaura = BDhanaura陀奴罗
	f.陀奴罗Dhanaura.SetParent(f)
	
	f.陀卢Dharau = BDharau陀卢
	f.陀卢Dharau.SetParent(f)
	
	f.乾补Kannpur = BKannpur乾补
	f.乾补Kannpur.SetParent(f)
	
	f.尸佉罗时Shikarji = BShikarji尸佉罗时
	f.尸佉罗时Shikarji.SetParent(f)
	
}
