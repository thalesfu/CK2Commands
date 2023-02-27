package kara_bogaz

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type Kara_bogazCounty interface {
    feud.County
    BAmandor阿曼多尔() 	feud.Barony
    BBekdash贝克达什() 	feud.Barony
    BGeksay格克赛() 	feud.Barony
    BKadzhan卡占() 	feud.Barony
    BKara_bogaz卡拉博加兹() 	feud.Barony
    BKaradzhari卡拉贾里() 	feud.Barony
    BSevervykh谢韦尔维赫() 	feud.Barony
}

type 卡拉博加兹Kara_bogazCounty struct {
	feud.BaseCounty
	阿曼多尔Amandor 	feud.Barony
	贝克达什Bekdash 	feud.Barony
	格克赛Geksay 	feud.Barony
	卡占Kadzhan 	feud.Barony
	卡拉博加兹Kara_bogaz 	feud.Barony
	卡拉贾里Karadzhari 	feud.Barony
	谢韦尔维赫Severvykh 	feud.Barony
}

func (c *卡拉博加兹Kara_bogazCounty) BAmandor阿曼多尔() feud.Barony {
	return c.阿曼多尔Amandor
}
    
func (c *卡拉博加兹Kara_bogazCounty) BBekdash贝克达什() feud.Barony {
	return c.贝克达什Bekdash
}
    
func (c *卡拉博加兹Kara_bogazCounty) BGeksay格克赛() feud.Barony {
	return c.格克赛Geksay
}
    
func (c *卡拉博加兹Kara_bogazCounty) BKadzhan卡占() feud.Barony {
	return c.卡占Kadzhan
}
    
func (c *卡拉博加兹Kara_bogazCounty) BKara_bogaz卡拉博加兹() feud.Barony {
	return c.卡拉博加兹Kara_bogaz
}
    
func (c *卡拉博加兹Kara_bogazCounty) BKaradzhari卡拉贾里() feud.Barony {
	return c.卡拉贾里Karadzhari
}
    
func (c *卡拉博加兹Kara_bogazCounty) BSevervykh谢韦尔维赫() feud.Barony {
	return c.谢韦尔维赫Severvykh
}
    
var CKara_bogaz卡拉博加兹 Kara_bogazCounty = &卡拉博加兹Kara_bogazCounty{}

func init() {
	f := CKara_bogaz卡拉博加兹.(*卡拉博加兹Kara_bogazCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1788",
		Title:     "kara_bogaz",
		TitleName: "卡拉博加兹",
		TitleCode: "c_kara_bogaz",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿曼多尔Amandor = BAmandor阿曼多尔
	f.阿曼多尔Amandor.SetParent(f)
	
	f.贝克达什Bekdash = BBekdash贝克达什
	f.贝克达什Bekdash.SetParent(f)
	
	f.格克赛Geksay = BGeksay格克赛
	f.格克赛Geksay.SetParent(f)
	
	f.卡占Kadzhan = BKadzhan卡占
	f.卡占Kadzhan.SetParent(f)
	
	f.卡拉博加兹Kara_bogaz = BKara_bogaz卡拉博加兹
	f.卡拉博加兹Kara_bogaz.SetParent(f)
	
	f.卡拉贾里Karadzhari = BKaradzhari卡拉贾里
	f.卡拉贾里Karadzhari.SetParent(f)
	
	f.谢韦尔维赫Severvykh = BSevervykh谢韦尔维赫
	f.谢韦尔维赫Severvykh.SetParent(f)
	
}
