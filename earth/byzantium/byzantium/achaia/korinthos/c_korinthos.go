package korinthos

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type KorinthosCounty interface {
    feud.County
    BArgos阿尔戈斯() 	feud.Barony
    BCorinth科林斯() 	feud.Barony
    BMegapoli巨城() 	feud.Barony
    BNauplion纳夫普利翁() 	feud.Barony
    BPassava帕萨瓦() 	feud.Barony
    BVeligosti韦利戈斯蒂() 	feud.Barony
    BVostitza沃斯蒂察() 	feud.Barony
    BZemenos泽梅诺斯() 	feud.Barony
}

type 科林斯KorinthosCounty struct {
	feud.BaseCounty
	阿尔戈斯Argos 	feud.Barony
	科林斯Corinth 	feud.Barony
	巨城Megapoli 	feud.Barony
	纳夫普利翁Nauplion 	feud.Barony
	帕萨瓦Passava 	feud.Barony
	韦利戈斯蒂Veligosti 	feud.Barony
	沃斯蒂察Vostitza 	feud.Barony
	泽梅诺斯Zemenos 	feud.Barony
}

func (c *科林斯KorinthosCounty) BArgos阿尔戈斯() feud.Barony {
	return c.阿尔戈斯Argos
}
    
func (c *科林斯KorinthosCounty) BCorinth科林斯() feud.Barony {
	return c.科林斯Corinth
}
    
func (c *科林斯KorinthosCounty) BMegapoli巨城() feud.Barony {
	return c.巨城Megapoli
}
    
func (c *科林斯KorinthosCounty) BNauplion纳夫普利翁() feud.Barony {
	return c.纳夫普利翁Nauplion
}
    
func (c *科林斯KorinthosCounty) BPassava帕萨瓦() feud.Barony {
	return c.帕萨瓦Passava
}
    
func (c *科林斯KorinthosCounty) BVeligosti韦利戈斯蒂() feud.Barony {
	return c.韦利戈斯蒂Veligosti
}
    
func (c *科林斯KorinthosCounty) BVostitza沃斯蒂察() feud.Barony {
	return c.沃斯蒂察Vostitza
}
    
func (c *科林斯KorinthosCounty) BZemenos泽梅诺斯() feud.Barony {
	return c.泽梅诺斯Zemenos
}
    
var CKorinthos科林斯 KorinthosCounty = &科林斯KorinthosCounty{}

func init() {
	f := CKorinthos科林斯.(*科林斯KorinthosCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "481",
		Title:     "korinthos",
		TitleName: "科林斯",
		TitleCode: "c_korinthos",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿尔戈斯Argos = BArgos阿尔戈斯
	f.阿尔戈斯Argos.SetParent(f)
	
	f.科林斯Corinth = BCorinth科林斯
	f.科林斯Corinth.SetParent(f)
	
	f.巨城Megapoli = BMegapoli巨城
	f.巨城Megapoli.SetParent(f)
	
	f.纳夫普利翁Nauplion = BNauplion纳夫普利翁
	f.纳夫普利翁Nauplion.SetParent(f)
	
	f.帕萨瓦Passava = BPassava帕萨瓦
	f.帕萨瓦Passava.SetParent(f)
	
	f.韦利戈斯蒂Veligosti = BVeligosti韦利戈斯蒂
	f.韦利戈斯蒂Veligosti.SetParent(f)
	
	f.沃斯蒂察Vostitza = BVostitza沃斯蒂察
	f.沃斯蒂察Vostitza.SetParent(f)
	
	f.泽梅诺斯Zemenos = BZemenos泽梅诺斯
	f.泽梅诺斯Zemenos.SetParent(f)
	
}
