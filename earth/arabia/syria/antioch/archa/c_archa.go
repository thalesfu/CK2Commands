package archa

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type ArchaCounty interface {
    feud.County
    BArcha阿尔卡() 	feud.Barony
    BChastelblanc布朗堡() 	feud.Barony
    BFamia阿法米亚() 	feud.Barony
    BKafroun卡夫伦() 	feud.Barony
    BMasyaf迈斯亚夫() 	feud.Barony
    BShayzar西扎() 	feud.Barony
    BTreiz特耶兹() 	feud.Barony
}

type 阿尔卡ArchaCounty struct {
	feud.BaseCounty
	阿尔卡Archa 	feud.Barony
	布朗堡Chastelblanc 	feud.Barony
	阿法米亚Famia 	feud.Barony
	卡夫伦Kafroun 	feud.Barony
	迈斯亚夫Masyaf 	feud.Barony
	西扎Shayzar 	feud.Barony
	特耶兹Treiz 	feud.Barony
}

func (c *阿尔卡ArchaCounty) BArcha阿尔卡() feud.Barony {
	return c.阿尔卡Archa
}
    
func (c *阿尔卡ArchaCounty) BChastelblanc布朗堡() feud.Barony {
	return c.布朗堡Chastelblanc
}
    
func (c *阿尔卡ArchaCounty) BFamia阿法米亚() feud.Barony {
	return c.阿法米亚Famia
}
    
func (c *阿尔卡ArchaCounty) BKafroun卡夫伦() feud.Barony {
	return c.卡夫伦Kafroun
}
    
func (c *阿尔卡ArchaCounty) BMasyaf迈斯亚夫() feud.Barony {
	return c.迈斯亚夫Masyaf
}
    
func (c *阿尔卡ArchaCounty) BShayzar西扎() feud.Barony {
	return c.西扎Shayzar
}
    
func (c *阿尔卡ArchaCounty) BTreiz特耶兹() feud.Barony {
	return c.特耶兹Treiz
}
    
var CArcha阿尔卡 ArchaCounty = &阿尔卡ArchaCounty{}

func init() {
	f := CArcha阿尔卡.(*阿尔卡ArchaCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "765",
		Title:     "archa",
		TitleName: "阿尔卡",
		TitleCode: "c_archa",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿尔卡Archa = BArcha阿尔卡
	f.阿尔卡Archa.SetParent(f)
	
	f.布朗堡Chastelblanc = BChastelblanc布朗堡
	f.布朗堡Chastelblanc.SetParent(f)
	
	f.阿法米亚Famia = BFamia阿法米亚
	f.阿法米亚Famia.SetParent(f)
	
	f.卡夫伦Kafroun = BKafroun卡夫伦
	f.卡夫伦Kafroun.SetParent(f)
	
	f.迈斯亚夫Masyaf = BMasyaf迈斯亚夫
	f.迈斯亚夫Masyaf.SetParent(f)
	
	f.西扎Shayzar = BShayzar西扎
	f.西扎Shayzar.SetParent(f)
	
	f.特耶兹Treiz = BTreiz特耶兹
	f.特耶兹Treiz.SetParent(f)
	
}
