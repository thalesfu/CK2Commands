package kherla

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type KherlaCounty interface {
    feud.County
    BAnjangaon安禅那伽罗摩() 	feud.Barony
    BGawilghar格维尔格尔() 	feud.Barony
    BKherla稽剌罗() 	feud.Barony
    BMauli摩梨() 	feud.Barony
    BPachmarhi伯杰默里() 	feud.Barony
    BPemayangtse贝玛扬泽() 	feud.Barony
    BPithorai毗豆罗() 	feud.Barony
}

type 稽剌罗KherlaCounty struct {
	feud.BaseCounty
	安禅那伽罗摩Anjangaon 	feud.Barony
	格维尔格尔Gawilghar 	feud.Barony
	稽剌罗Kherla 	feud.Barony
	摩梨Mauli 	feud.Barony
	伯杰默里Pachmarhi 	feud.Barony
	贝玛扬泽Pemayangtse 	feud.Barony
	毗豆罗Pithorai 	feud.Barony
}

func (c *稽剌罗KherlaCounty) BAnjangaon安禅那伽罗摩() feud.Barony {
	return c.安禅那伽罗摩Anjangaon
}
    
func (c *稽剌罗KherlaCounty) BGawilghar格维尔格尔() feud.Barony {
	return c.格维尔格尔Gawilghar
}
    
func (c *稽剌罗KherlaCounty) BKherla稽剌罗() feud.Barony {
	return c.稽剌罗Kherla
}
    
func (c *稽剌罗KherlaCounty) BMauli摩梨() feud.Barony {
	return c.摩梨Mauli
}
    
func (c *稽剌罗KherlaCounty) BPachmarhi伯杰默里() feud.Barony {
	return c.伯杰默里Pachmarhi
}
    
func (c *稽剌罗KherlaCounty) BPemayangtse贝玛扬泽() feud.Barony {
	return c.贝玛扬泽Pemayangtse
}
    
func (c *稽剌罗KherlaCounty) BPithorai毗豆罗() feud.Barony {
	return c.毗豆罗Pithorai
}
    
var CKherla稽剌罗 KherlaCounty = &稽剌罗KherlaCounty{}

func init() {
	f := CKherla稽剌罗.(*稽剌罗KherlaCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1286",
		Title:     "kherla",
		TitleName: "稽剌罗",
		TitleCode: "c_kherla",
		Baronies:  map[string]feud.Barony{},
	}

	f.安禅那伽罗摩Anjangaon = BAnjangaon安禅那伽罗摩
	f.安禅那伽罗摩Anjangaon.SetParent(f)
	
	f.格维尔格尔Gawilghar = BGawilghar格维尔格尔
	f.格维尔格尔Gawilghar.SetParent(f)
	
	f.稽剌罗Kherla = BKherla稽剌罗
	f.稽剌罗Kherla.SetParent(f)
	
	f.摩梨Mauli = BMauli摩梨
	f.摩梨Mauli.SetParent(f)
	
	f.伯杰默里Pachmarhi = BPachmarhi伯杰默里
	f.伯杰默里Pachmarhi.SetParent(f)
	
	f.贝玛扬泽Pemayangtse = BPemayangtse贝玛扬泽
	f.贝玛扬泽Pemayangtse.SetParent(f)
	
	f.毗豆罗Pithorai = BPithorai毗豆罗
	f.毗豆罗Pithorai.SetParent(f)
	
}
