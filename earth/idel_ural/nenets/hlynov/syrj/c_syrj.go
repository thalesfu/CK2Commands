package syrj

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type SyrjCounty interface {
    feud.County
    BEmva叶姆瓦() 	feud.Barony
    BEzhva埃日瓦() 	feud.Barony
    BMikun米昆() 	feud.Barony
    BPyras佩拉斯() 	feud.Barony
    BSindor辛多尔() 	feud.Barony
    BYugydyag尤格德亚格() 	feud.Barony
    BZheshart热沙尔特() 	feud.Barony
}

type 西里SyrjCounty struct {
	feud.BaseCounty
	叶姆瓦Emva 	feud.Barony
	埃日瓦Ezhva 	feud.Barony
	米昆Mikun 	feud.Barony
	佩拉斯Pyras 	feud.Barony
	辛多尔Sindor 	feud.Barony
	尤格德亚格Yugydyag 	feud.Barony
	热沙尔特Zheshart 	feud.Barony
}

func (c *西里SyrjCounty) BEmva叶姆瓦() feud.Barony {
	return c.叶姆瓦Emva
}
    
func (c *西里SyrjCounty) BEzhva埃日瓦() feud.Barony {
	return c.埃日瓦Ezhva
}
    
func (c *西里SyrjCounty) BMikun米昆() feud.Barony {
	return c.米昆Mikun
}
    
func (c *西里SyrjCounty) BPyras佩拉斯() feud.Barony {
	return c.佩拉斯Pyras
}
    
func (c *西里SyrjCounty) BSindor辛多尔() feud.Barony {
	return c.辛多尔Sindor
}
    
func (c *西里SyrjCounty) BYugydyag尤格德亚格() feud.Barony {
	return c.尤格德亚格Yugydyag
}
    
func (c *西里SyrjCounty) BZheshart热沙尔特() feud.Barony {
	return c.热沙尔特Zheshart
}
    
var CSyrj西里 SyrjCounty = &西里SyrjCounty{}

func init() {
	f := CSyrj西里.(*西里SyrjCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "399",
		Title:     "syrj",
		TitleName: "西里",
		TitleCode: "c_syrj",
		Baronies:  map[string]feud.Barony{},
	}

	f.叶姆瓦Emva = BEmva叶姆瓦
	f.叶姆瓦Emva.SetParent(f)
	
	f.埃日瓦Ezhva = BEzhva埃日瓦
	f.埃日瓦Ezhva.SetParent(f)
	
	f.米昆Mikun = BMikun米昆
	f.米昆Mikun.SetParent(f)
	
	f.佩拉斯Pyras = BPyras佩拉斯
	f.佩拉斯Pyras.SetParent(f)
	
	f.辛多尔Sindor = BSindor辛多尔
	f.辛多尔Sindor.SetParent(f)
	
	f.尤格德亚格Yugydyag = BYugydyag尤格德亚格
	f.尤格德亚格Yugydyag.SetParent(f)
	
	f.热沙尔特Zheshart = BZheshart热沙尔特
	f.热沙尔特Zheshart.SetParent(f)
	
}
