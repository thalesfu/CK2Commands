package kimak

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type KimakCounty interface {
    feud.County
    BArkalyk阿尔卡雷克() 	feud.Barony
    BAtasu阿塔苏() 	feud.Barony
    BJutasu珠塔苏() 	feud.Barony
    BKarasor卡拉索尔() 	feud.Barony
    BOpak奥帕克() 	feud.Barony
    BTengiz田吉兹() 	feud.Barony
    BTenul捷努尔() 	feud.Barony
}

type 卡拉索尔KimakCounty struct {
	feud.BaseCounty
	阿尔卡雷克Arkalyk 	feud.Barony
	阿塔苏Atasu 	feud.Barony
	珠塔苏Jutasu 	feud.Barony
	卡拉索尔Karasor 	feud.Barony
	奥帕克Opak 	feud.Barony
	田吉兹Tengiz 	feud.Barony
	捷努尔Tenul 	feud.Barony
}

func (c *卡拉索尔KimakCounty) BArkalyk阿尔卡雷克() feud.Barony {
	return c.阿尔卡雷克Arkalyk
}
    
func (c *卡拉索尔KimakCounty) BAtasu阿塔苏() feud.Barony {
	return c.阿塔苏Atasu
}
    
func (c *卡拉索尔KimakCounty) BJutasu珠塔苏() feud.Barony {
	return c.珠塔苏Jutasu
}
    
func (c *卡拉索尔KimakCounty) BKarasor卡拉索尔() feud.Barony {
	return c.卡拉索尔Karasor
}
    
func (c *卡拉索尔KimakCounty) BOpak奥帕克() feud.Barony {
	return c.奥帕克Opak
}
    
func (c *卡拉索尔KimakCounty) BTengiz田吉兹() feud.Barony {
	return c.田吉兹Tengiz
}
    
func (c *卡拉索尔KimakCounty) BTenul捷努尔() feud.Barony {
	return c.捷努尔Tenul
}
    
var CKimak卡拉索尔 KimakCounty = &卡拉索尔KimakCounty{}

func init() {
	f := CKimak卡拉索尔.(*卡拉索尔KimakCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1435",
		Title:     "kimak",
		TitleName: "卡拉索尔",
		TitleCode: "c_kimak",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿尔卡雷克Arkalyk = BArkalyk阿尔卡雷克
	f.阿尔卡雷克Arkalyk.SetParent(f)
	
	f.阿塔苏Atasu = BAtasu阿塔苏
	f.阿塔苏Atasu.SetParent(f)
	
	f.珠塔苏Jutasu = BJutasu珠塔苏
	f.珠塔苏Jutasu.SetParent(f)
	
	f.卡拉索尔Karasor = BKarasor卡拉索尔
	f.卡拉索尔Karasor.SetParent(f)
	
	f.奥帕克Opak = BOpak奥帕克
	f.奥帕克Opak.SetParent(f)
	
	f.田吉兹Tengiz = BTengiz田吉兹
	f.田吉兹Tengiz.SetParent(f)
	
	f.捷努尔Tenul = BTenul捷努尔
	f.捷努尔Tenul.SetParent(f)
	
}
