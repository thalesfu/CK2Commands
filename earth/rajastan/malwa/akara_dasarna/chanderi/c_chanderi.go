package chanderi

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type ChanderiCounty interface {
    feud.County
    BAharji阿诃吉() 	feud.Barony
    BChanderi旃提梨() 	feud.Barony
    BGarh_kundar迦尔贡陀尔() 	feud.Barony
    BJarai_ka_math奢罗迦摩斯() 	feud.Barony
    BJhansi占西() 	feud.Barony
    BKhurai库赖() 	feud.Barony
    BLuacchagiri捋车耆厘() 	feud.Barony
}

type 旃提梨ChanderiCounty struct {
	feud.BaseCounty
	阿诃吉Aharji 	feud.Barony
	旃提梨Chanderi 	feud.Barony
	迦尔贡陀尔Garh_kundar 	feud.Barony
	奢罗迦摩斯Jarai_ka_math 	feud.Barony
	占西Jhansi 	feud.Barony
	库赖Khurai 	feud.Barony
	捋车耆厘Luacchagiri 	feud.Barony
}

func (c *旃提梨ChanderiCounty) BAharji阿诃吉() feud.Barony {
	return c.阿诃吉Aharji
}
    
func (c *旃提梨ChanderiCounty) BChanderi旃提梨() feud.Barony {
	return c.旃提梨Chanderi
}
    
func (c *旃提梨ChanderiCounty) BGarh_kundar迦尔贡陀尔() feud.Barony {
	return c.迦尔贡陀尔Garh_kundar
}
    
func (c *旃提梨ChanderiCounty) BJarai_ka_math奢罗迦摩斯() feud.Barony {
	return c.奢罗迦摩斯Jarai_ka_math
}
    
func (c *旃提梨ChanderiCounty) BJhansi占西() feud.Barony {
	return c.占西Jhansi
}
    
func (c *旃提梨ChanderiCounty) BKhurai库赖() feud.Barony {
	return c.库赖Khurai
}
    
func (c *旃提梨ChanderiCounty) BLuacchagiri捋车耆厘() feud.Barony {
	return c.捋车耆厘Luacchagiri
}
    
var CChanderi旃提梨 ChanderiCounty = &旃提梨ChanderiCounty{}

func init() {
	f := CChanderi旃提梨.(*旃提梨ChanderiCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1298",
		Title:     "chanderi",
		TitleName: "旃提梨",
		TitleCode: "c_chanderi",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿诃吉Aharji = BAharji阿诃吉
	f.阿诃吉Aharji.SetParent(f)
	
	f.旃提梨Chanderi = BChanderi旃提梨
	f.旃提梨Chanderi.SetParent(f)
	
	f.迦尔贡陀尔Garh_kundar = BGarh_kundar迦尔贡陀尔
	f.迦尔贡陀尔Garh_kundar.SetParent(f)
	
	f.奢罗迦摩斯Jarai_ka_math = BJarai_ka_math奢罗迦摩斯
	f.奢罗迦摩斯Jarai_ka_math.SetParent(f)
	
	f.占西Jhansi = BJhansi占西
	f.占西Jhansi.SetParent(f)
	
	f.库赖Khurai = BKhurai库赖
	f.库赖Khurai.SetParent(f)
	
	f.捋车耆厘Luacchagiri = BLuacchagiri捋车耆厘
	f.捋车耆厘Luacchagiri.SetParent(f)
	
}
