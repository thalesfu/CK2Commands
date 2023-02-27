package kazakh

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type KazakhCounty interface {
    feud.County
    BJezkazgan杰兹卡兹干() 	feud.Barony
    BKarsakpay卡尔萨克帕伊() 	feud.Barony
    BKengir肯吉尔() 	feud.Barony
    BZhairem扎伊列姆() 	feud.Barony
    BZhezdi杰兹迪() 	feud.Barony
    BZhezgir杰孜吉尔() 	feud.Barony
    BZhezkazgan热兹卡兹甘() 	feud.Barony
}

type 哈萨克KazakhCounty struct {
	feud.BaseCounty
	杰兹卡兹干Jezkazgan 	feud.Barony
	卡尔萨克帕伊Karsakpay 	feud.Barony
	肯吉尔Kengir 	feud.Barony
	扎伊列姆Zhairem 	feud.Barony
	杰兹迪Zhezdi 	feud.Barony
	杰孜吉尔Zhezgir 	feud.Barony
	热兹卡兹甘Zhezkazgan 	feud.Barony
}

func (c *哈萨克KazakhCounty) BJezkazgan杰兹卡兹干() feud.Barony {
	return c.杰兹卡兹干Jezkazgan
}
    
func (c *哈萨克KazakhCounty) BKarsakpay卡尔萨克帕伊() feud.Barony {
	return c.卡尔萨克帕伊Karsakpay
}
    
func (c *哈萨克KazakhCounty) BKengir肯吉尔() feud.Barony {
	return c.肯吉尔Kengir
}
    
func (c *哈萨克KazakhCounty) BZhairem扎伊列姆() feud.Barony {
	return c.扎伊列姆Zhairem
}
    
func (c *哈萨克KazakhCounty) BZhezdi杰兹迪() feud.Barony {
	return c.杰兹迪Zhezdi
}
    
func (c *哈萨克KazakhCounty) BZhezgir杰孜吉尔() feud.Barony {
	return c.杰孜吉尔Zhezgir
}
    
func (c *哈萨克KazakhCounty) BZhezkazgan热兹卡兹甘() feud.Barony {
	return c.热兹卡兹甘Zhezkazgan
}
    
var CKazakh哈萨克 KazakhCounty = &哈萨克KazakhCounty{}

func init() {
	f := CKazakh哈萨克.(*哈萨克KazakhCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1433",
		Title:     "kazakh",
		TitleName: "哈萨克",
		TitleCode: "c_kazakh",
		Baronies:  map[string]feud.Barony{},
	}

	f.杰兹卡兹干Jezkazgan = BJezkazgan杰兹卡兹干
	f.杰兹卡兹干Jezkazgan.SetParent(f)
	
	f.卡尔萨克帕伊Karsakpay = BKarsakpay卡尔萨克帕伊
	f.卡尔萨克帕伊Karsakpay.SetParent(f)
	
	f.肯吉尔Kengir = BKengir肯吉尔
	f.肯吉尔Kengir.SetParent(f)
	
	f.扎伊列姆Zhairem = BZhairem扎伊列姆
	f.扎伊列姆Zhairem.SetParent(f)
	
	f.杰兹迪Zhezdi = BZhezdi杰兹迪
	f.杰兹迪Zhezdi.SetParent(f)
	
	f.杰孜吉尔Zhezgir = BZhezgir杰孜吉尔
	f.杰孜吉尔Zhezgir.SetParent(f)
	
	f.热兹卡兹甘Zhezkazgan = BZhezkazgan热兹卡兹甘
	f.热兹卡兹甘Zhezkazgan.SetParent(f)
	
}
