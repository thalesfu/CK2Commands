package suvarnagram

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type SuvarnagramCounty interface {
    feud.County
    BBokainagar菩迦伊城() 	feud.Barony
    BLakshmipuro罗迦施弥补罗() 	feud.Barony
    BMahendragarh摩醯因陀罗姞利呬() 	feud.Barony
    BNarikella那利蓟罗() 	feud.Barony
    BSatkhira萨吉罗() 	feud.Barony
    BSusanga苏僧伽() 	feud.Barony
    BSuvarnagram苏伐剌拿伽罗摩() 	feud.Barony
}

type 苏伐剌拿伽罗摩SuvarnagramCounty struct {
	feud.BaseCounty
	菩迦伊城Bokainagar 	feud.Barony
	罗迦施弥补罗Lakshmipuro 	feud.Barony
	摩醯因陀罗姞利呬Mahendragarh 	feud.Barony
	那利蓟罗Narikella 	feud.Barony
	萨吉罗Satkhira 	feud.Barony
	苏僧伽Susanga 	feud.Barony
	苏伐剌拿伽罗摩Suvarnagram 	feud.Barony
}

func (c *苏伐剌拿伽罗摩SuvarnagramCounty) BBokainagar菩迦伊城() feud.Barony {
	return c.菩迦伊城Bokainagar
}
    
func (c *苏伐剌拿伽罗摩SuvarnagramCounty) BLakshmipuro罗迦施弥补罗() feud.Barony {
	return c.罗迦施弥补罗Lakshmipuro
}
    
func (c *苏伐剌拿伽罗摩SuvarnagramCounty) BMahendragarh摩醯因陀罗姞利呬() feud.Barony {
	return c.摩醯因陀罗姞利呬Mahendragarh
}
    
func (c *苏伐剌拿伽罗摩SuvarnagramCounty) BNarikella那利蓟罗() feud.Barony {
	return c.那利蓟罗Narikella
}
    
func (c *苏伐剌拿伽罗摩SuvarnagramCounty) BSatkhira萨吉罗() feud.Barony {
	return c.萨吉罗Satkhira
}
    
func (c *苏伐剌拿伽罗摩SuvarnagramCounty) BSusanga苏僧伽() feud.Barony {
	return c.苏僧伽Susanga
}
    
func (c *苏伐剌拿伽罗摩SuvarnagramCounty) BSuvarnagram苏伐剌拿伽罗摩() feud.Barony {
	return c.苏伐剌拿伽罗摩Suvarnagram
}
    
var CSuvarnagram苏伐剌拿伽罗摩 SuvarnagramCounty = &苏伐剌拿伽罗摩SuvarnagramCounty{}

func init() {
	f := CSuvarnagram苏伐剌拿伽罗摩.(*苏伐剌拿伽罗摩SuvarnagramCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1324",
		Title:     "suvarnagram",
		TitleName: "苏伐剌拿伽罗摩",
		TitleCode: "c_suvarnagram",
		Baronies:  map[string]feud.Barony{},
	}

	f.菩迦伊城Bokainagar = BBokainagar菩迦伊城
	f.菩迦伊城Bokainagar.SetParent(f)
	
	f.罗迦施弥补罗Lakshmipuro = BLakshmipuro罗迦施弥补罗
	f.罗迦施弥补罗Lakshmipuro.SetParent(f)
	
	f.摩醯因陀罗姞利呬Mahendragarh = BMahendragarh摩醯因陀罗姞利呬
	f.摩醯因陀罗姞利呬Mahendragarh.SetParent(f)
	
	f.那利蓟罗Narikella = BNarikella那利蓟罗
	f.那利蓟罗Narikella.SetParent(f)
	
	f.萨吉罗Satkhira = BSatkhira萨吉罗
	f.萨吉罗Satkhira.SetParent(f)
	
	f.苏僧伽Susanga = BSusanga苏僧伽
	f.苏僧伽Susanga.SetParent(f)
	
	f.苏伐剌拿伽罗摩Suvarnagram = BSuvarnagram苏伐剌拿伽罗摩
	f.苏伐剌拿伽罗摩Suvarnagram.SetParent(f)
	
}
