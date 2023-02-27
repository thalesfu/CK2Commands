package kremenchuk

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type KremenchukCounty interface {
    feud.County
    BAlbay阿尔拜() 	feud.Barony
    BKatmysh卡特梅什() 	feud.Barony
    BMamadich玛玛迪奇() 	feud.Barony
    BOtarka奥塔尔卡() 	feud.Barony
    BSharan沙兰() 	feud.Barony
    BShumbut顺布特() 	feud.Barony
    BVolgakremenchuk克列缅丘格() 	feud.Barony
}

type 克列缅丘格KremenchukCounty struct {
	feud.BaseCounty
	阿尔拜Albay 	feud.Barony
	卡特梅什Katmysh 	feud.Barony
	玛玛迪奇Mamadich 	feud.Barony
	奥塔尔卡Otarka 	feud.Barony
	沙兰Sharan 	feud.Barony
	顺布特Shumbut 	feud.Barony
	克列缅丘格Volgakremenchuk 	feud.Barony
}

func (c *克列缅丘格KremenchukCounty) BAlbay阿尔拜() feud.Barony {
	return c.阿尔拜Albay
}
    
func (c *克列缅丘格KremenchukCounty) BKatmysh卡特梅什() feud.Barony {
	return c.卡特梅什Katmysh
}
    
func (c *克列缅丘格KremenchukCounty) BMamadich玛玛迪奇() feud.Barony {
	return c.玛玛迪奇Mamadich
}
    
func (c *克列缅丘格KremenchukCounty) BOtarka奥塔尔卡() feud.Barony {
	return c.奥塔尔卡Otarka
}
    
func (c *克列缅丘格KremenchukCounty) BSharan沙兰() feud.Barony {
	return c.沙兰Sharan
}
    
func (c *克列缅丘格KremenchukCounty) BShumbut顺布特() feud.Barony {
	return c.顺布特Shumbut
}
    
func (c *克列缅丘格KremenchukCounty) BVolgakremenchuk克列缅丘格() feud.Barony {
	return c.克列缅丘格Volgakremenchuk
}
    
var CKremenchuk克列缅丘格 KremenchukCounty = &克列缅丘格KremenchukCounty{}

func init() {
	f := CKremenchuk克列缅丘格.(*克列缅丘格KremenchukCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1716",
		Title:     "kremenchuk",
		TitleName: "克列缅丘格",
		TitleCode: "c_kremenchuk",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿尔拜Albay = BAlbay阿尔拜
	f.阿尔拜Albay.SetParent(f)
	
	f.卡特梅什Katmysh = BKatmysh卡特梅什
	f.卡特梅什Katmysh.SetParent(f)
	
	f.玛玛迪奇Mamadich = BMamadich玛玛迪奇
	f.玛玛迪奇Mamadich.SetParent(f)
	
	f.奥塔尔卡Otarka = BOtarka奥塔尔卡
	f.奥塔尔卡Otarka.SetParent(f)
	
	f.沙兰Sharan = BSharan沙兰
	f.沙兰Sharan.SetParent(f)
	
	f.顺布特Shumbut = BShumbut顺布特
	f.顺布特Shumbut.SetParent(f)
	
	f.克列缅丘格Volgakremenchuk = BVolgakremenchuk克列缅丘格
	f.克列缅丘格Volgakremenchuk.SetParent(f)
	
}
