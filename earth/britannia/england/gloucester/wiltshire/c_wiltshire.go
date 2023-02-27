package wiltshire

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type WiltshireCounty interface {
    feud.County
    BClarendon克拉伦登() 	feud.Barony
    BDevizes迪韦齐斯() 	feud.Barony
    BMalmesbury马姆斯伯里() 	feud.Barony
    BMarlborough莫尔伯勒() 	feud.Barony
    BRamsbury拉姆斯伯里() 	feud.Barony
    BSalisbury索尔兹伯里() 	feud.Barony
    BSarum塞勒姆() 	feud.Barony
    BWilton威尔顿() 	feud.Barony
}

type 威尔特郡WiltshireCounty struct {
	feud.BaseCounty
	克拉伦登Clarendon 	feud.Barony
	迪韦齐斯Devizes 	feud.Barony
	马姆斯伯里Malmesbury 	feud.Barony
	莫尔伯勒Marlborough 	feud.Barony
	拉姆斯伯里Ramsbury 	feud.Barony
	索尔兹伯里Salisbury 	feud.Barony
	塞勒姆Sarum 	feud.Barony
	威尔顿Wilton 	feud.Barony
}

func (c *威尔特郡WiltshireCounty) BClarendon克拉伦登() feud.Barony {
	return c.克拉伦登Clarendon
}
    
func (c *威尔特郡WiltshireCounty) BDevizes迪韦齐斯() feud.Barony {
	return c.迪韦齐斯Devizes
}
    
func (c *威尔特郡WiltshireCounty) BMalmesbury马姆斯伯里() feud.Barony {
	return c.马姆斯伯里Malmesbury
}
    
func (c *威尔特郡WiltshireCounty) BMarlborough莫尔伯勒() feud.Barony {
	return c.莫尔伯勒Marlborough
}
    
func (c *威尔特郡WiltshireCounty) BRamsbury拉姆斯伯里() feud.Barony {
	return c.拉姆斯伯里Ramsbury
}
    
func (c *威尔特郡WiltshireCounty) BSalisbury索尔兹伯里() feud.Barony {
	return c.索尔兹伯里Salisbury
}
    
func (c *威尔特郡WiltshireCounty) BSarum塞勒姆() feud.Barony {
	return c.塞勒姆Sarum
}
    
func (c *威尔特郡WiltshireCounty) BWilton威尔顿() feud.Barony {
	return c.威尔顿Wilton
}
    
var CWiltshire威尔特郡 WiltshireCounty = &威尔特郡WiltshireCounty{}

func init() {
	f := CWiltshire威尔特郡.(*威尔特郡WiltshireCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "23",
		Title:     "wiltshire",
		TitleName: "威尔特郡",
		TitleCode: "c_wiltshire",
		Baronies:  map[string]feud.Barony{},
	}

	f.克拉伦登Clarendon = BClarendon克拉伦登
	f.克拉伦登Clarendon.SetParent(f)
	
	f.迪韦齐斯Devizes = BDevizes迪韦齐斯
	f.迪韦齐斯Devizes.SetParent(f)
	
	f.马姆斯伯里Malmesbury = BMalmesbury马姆斯伯里
	f.马姆斯伯里Malmesbury.SetParent(f)
	
	f.莫尔伯勒Marlborough = BMarlborough莫尔伯勒
	f.莫尔伯勒Marlborough.SetParent(f)
	
	f.拉姆斯伯里Ramsbury = BRamsbury拉姆斯伯里
	f.拉姆斯伯里Ramsbury.SetParent(f)
	
	f.索尔兹伯里Salisbury = BSalisbury索尔兹伯里
	f.索尔兹伯里Salisbury.SetParent(f)
	
	f.塞勒姆Sarum = BSarum塞勒姆
	f.塞勒姆Sarum.SetParent(f)
	
	f.威尔顿Wilton = BWilton威尔顿
	f.威尔顿Wilton.SetParent(f)
	
}
