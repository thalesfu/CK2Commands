package amdo

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type AmdoCounty interface {
    feud.County
    BAmdo安多() 	feud.Barony
    BCona错那() 	feud.Barony
    BGangnyi岗尼() 	feud.Barony
    BNyainrong聂荣() 	feud.Barony
    BQangma强玛() 	feud.Barony
    BSewa色务() 	feud.Barony
    BZaring扎仁() 	feud.Barony
}

type 安多AmdoCounty struct {
	feud.BaseCounty
	安多Amdo 	feud.Barony
	错那Cona 	feud.Barony
	岗尼Gangnyi 	feud.Barony
	聂荣Nyainrong 	feud.Barony
	强玛Qangma 	feud.Barony
	色务Sewa 	feud.Barony
	扎仁Zaring 	feud.Barony
}

func (c *安多AmdoCounty) BAmdo安多() feud.Barony {
	return c.安多Amdo
}
    
func (c *安多AmdoCounty) BCona错那() feud.Barony {
	return c.错那Cona
}
    
func (c *安多AmdoCounty) BGangnyi岗尼() feud.Barony {
	return c.岗尼Gangnyi
}
    
func (c *安多AmdoCounty) BNyainrong聂荣() feud.Barony {
	return c.聂荣Nyainrong
}
    
func (c *安多AmdoCounty) BQangma强玛() feud.Barony {
	return c.强玛Qangma
}
    
func (c *安多AmdoCounty) BSewa色务() feud.Barony {
	return c.色务Sewa
}
    
func (c *安多AmdoCounty) BZaring扎仁() feud.Barony {
	return c.扎仁Zaring
}
    
var CAmdo安多 AmdoCounty = &安多AmdoCounty{}

func init() {
	f := CAmdo安多.(*安多AmdoCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1558",
		Title:     "amdo",
		TitleName: "安多",
		TitleCode: "c_amdo",
		Baronies:  map[string]feud.Barony{},
	}

	f.安多Amdo = BAmdo安多
	f.安多Amdo.SetParent(f)
	
	f.错那Cona = BCona错那
	f.错那Cona.SetParent(f)
	
	f.岗尼Gangnyi = BGangnyi岗尼
	f.岗尼Gangnyi.SetParent(f)
	
	f.聂荣Nyainrong = BNyainrong聂荣
	f.聂荣Nyainrong.SetParent(f)
	
	f.强玛Qangma = BQangma强玛
	f.强玛Qangma.SetParent(f)
	
	f.色务Sewa = BSewa色务
	f.色务Sewa.SetParent(f)
	
	f.扎仁Zaring = BZaring扎仁
	f.扎仁Zaring.SetParent(f)
	
}
