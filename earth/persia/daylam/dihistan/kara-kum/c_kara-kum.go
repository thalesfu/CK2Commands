package kara-kum

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type Kara-kumCounty interface {
    feud.County
    BFarava法拉瓦() 	feud.Barony
    BGazanjyk卡赞吉克() 	feud.Barony
    BKhodzhakala霍贾卡拉() 	feud.Barony
    BKochevka科切夫卡() 	feud.Barony
    BKyzyl_arvat克孜勒阿尔瓦特() 	feud.Barony
    BOhk奥赫() 	feud.Barony
    BToutli托特利() 	feud.Barony
}

type 法拉瓦Kara-kumCounty struct {
	feud.BaseCounty
	法拉瓦Farava 	feud.Barony
	卡赞吉克Gazanjyk 	feud.Barony
	霍贾卡拉Khodzhakala 	feud.Barony
	科切夫卡Kochevka 	feud.Barony
	克孜勒阿尔瓦特Kyzyl_arvat 	feud.Barony
	奥赫Ohk 	feud.Barony
	托特利Toutli 	feud.Barony
}

func (c *法拉瓦Kara-kumCounty) BFarava法拉瓦() feud.Barony {
	return c.法拉瓦Farava
}
    
func (c *法拉瓦Kara-kumCounty) BGazanjyk卡赞吉克() feud.Barony {
	return c.卡赞吉克Gazanjyk
}
    
func (c *法拉瓦Kara-kumCounty) BKhodzhakala霍贾卡拉() feud.Barony {
	return c.霍贾卡拉Khodzhakala
}
    
func (c *法拉瓦Kara-kumCounty) BKochevka科切夫卡() feud.Barony {
	return c.科切夫卡Kochevka
}
    
func (c *法拉瓦Kara-kumCounty) BKyzyl_arvat克孜勒阿尔瓦特() feud.Barony {
	return c.克孜勒阿尔瓦特Kyzyl_arvat
}
    
func (c *法拉瓦Kara-kumCounty) BOhk奥赫() feud.Barony {
	return c.奥赫Ohk
}
    
func (c *法拉瓦Kara-kumCounty) BToutli托特利() feud.Barony {
	return c.托特利Toutli
}
    
var CKara-kum法拉瓦 Kara-kumCounty = &法拉瓦Kara-kumCounty{}

func init() {
	f := CKara-kum法拉瓦.(*法拉瓦Kara-kumCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "627",
		Title:     "kara-kum",
		TitleName: "法拉瓦",
		TitleCode: "c_kara-kum",
		Baronies:  map[string]feud.Barony{},
	}

	f.法拉瓦Farava = BFarava法拉瓦
	f.法拉瓦Farava.SetParent(f)
	
	f.卡赞吉克Gazanjyk = BGazanjyk卡赞吉克
	f.卡赞吉克Gazanjyk.SetParent(f)
	
	f.霍贾卡拉Khodzhakala = BKhodzhakala霍贾卡拉
	f.霍贾卡拉Khodzhakala.SetParent(f)
	
	f.科切夫卡Kochevka = BKochevka科切夫卡
	f.科切夫卡Kochevka.SetParent(f)
	
	f.克孜勒阿尔瓦特Kyzyl_arvat = BKyzyl_arvat克孜勒阿尔瓦特
	f.克孜勒阿尔瓦特Kyzyl_arvat.SetParent(f)
	
	f.奥赫Ohk = BOhk奥赫
	f.奥赫Ohk.SetParent(f)
	
	f.托特利Toutli = BToutli托特利
	f.托特利Toutli.SetParent(f)
	
}
