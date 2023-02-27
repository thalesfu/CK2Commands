package smyrna

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type SmyrnaCounty interface {
    feud.County
    BAdramyttion阿德拉密提翁() 	feud.Barony
    BChio希俄() 	feud.Barony
    BErythrai埃利特莱() 	feud.Barony
    BKlazomeanai克莱索美奈() 	feud.Barony
    BKydonia基多尼亚() 	feud.Barony
    BPergamon帕加马() 	feud.Barony
    BPhokaia菲莱雅() 	feud.Barony
    BSmyrna士麦拿() 	feud.Barony
}

type 士麦拿SmyrnaCounty struct {
	feud.BaseCounty
	阿德拉密提翁Adramyttion 	feud.Barony
	希俄Chio 	feud.Barony
	埃利特莱Erythrai 	feud.Barony
	克莱索美奈Klazomeanai 	feud.Barony
	基多尼亚Kydonia 	feud.Barony
	帕加马Pergamon 	feud.Barony
	菲莱雅Phokaia 	feud.Barony
	士麦拿Smyrna 	feud.Barony
}

func (c *士麦拿SmyrnaCounty) BAdramyttion阿德拉密提翁() feud.Barony {
	return c.阿德拉密提翁Adramyttion
}
    
func (c *士麦拿SmyrnaCounty) BChio希俄() feud.Barony {
	return c.希俄Chio
}
    
func (c *士麦拿SmyrnaCounty) BErythrai埃利特莱() feud.Barony {
	return c.埃利特莱Erythrai
}
    
func (c *士麦拿SmyrnaCounty) BKlazomeanai克莱索美奈() feud.Barony {
	return c.克莱索美奈Klazomeanai
}
    
func (c *士麦拿SmyrnaCounty) BKydonia基多尼亚() feud.Barony {
	return c.基多尼亚Kydonia
}
    
func (c *士麦拿SmyrnaCounty) BPergamon帕加马() feud.Barony {
	return c.帕加马Pergamon
}
    
func (c *士麦拿SmyrnaCounty) BPhokaia菲莱雅() feud.Barony {
	return c.菲莱雅Phokaia
}
    
func (c *士麦拿SmyrnaCounty) BSmyrna士麦拿() feud.Barony {
	return c.士麦拿Smyrna
}
    
var CSmyrna士麦拿 SmyrnaCounty = &士麦拿SmyrnaCounty{}

func init() {
	f := CSmyrna士麦拿.(*士麦拿SmyrnaCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "745",
		Title:     "smyrna",
		TitleName: "士麦拿",
		TitleCode: "c_smyrna",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿德拉密提翁Adramyttion = BAdramyttion阿德拉密提翁
	f.阿德拉密提翁Adramyttion.SetParent(f)
	
	f.希俄Chio = BChio希俄
	f.希俄Chio.SetParent(f)
	
	f.埃利特莱Erythrai = BErythrai埃利特莱
	f.埃利特莱Erythrai.SetParent(f)
	
	f.克莱索美奈Klazomeanai = BKlazomeanai克莱索美奈
	f.克莱索美奈Klazomeanai.SetParent(f)
	
	f.基多尼亚Kydonia = BKydonia基多尼亚
	f.基多尼亚Kydonia.SetParent(f)
	
	f.帕加马Pergamon = BPergamon帕加马
	f.帕加马Pergamon.SetParent(f)
	
	f.菲莱雅Phokaia = BPhokaia菲莱雅
	f.菲莱雅Phokaia.SetParent(f)
	
	f.士麦拿Smyrna = BSmyrna士麦拿
	f.士麦拿Smyrna.SetParent(f)
	
}
