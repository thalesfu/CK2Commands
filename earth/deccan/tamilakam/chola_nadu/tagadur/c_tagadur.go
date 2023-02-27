package tagadur

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type TagadurCounty interface {
    feud.County
    BJinji旖耆() 	feud.Barony
    BPadaividu波多毗都() 	feud.Barony
    BSatyamangalam娑底也瞢伽罗() 	feud.Barony
    BSiyyamangalam斯耶门伽愣() 	feud.Barony
    BTagadur多伽头罗() 	feud.Barony
    BTirukkovalur提卢伽婆卢() 	feud.Barony
    BTiruvannamalai提卢瓦那摩罗() 	feud.Barony
}

type 多伽头罗TagadurCounty struct {
	feud.BaseCounty
	旖耆Jinji 	feud.Barony
	波多毗都Padaividu 	feud.Barony
	娑底也瞢伽罗Satyamangalam 	feud.Barony
	斯耶门伽愣Siyyamangalam 	feud.Barony
	多伽头罗Tagadur 	feud.Barony
	提卢伽婆卢Tirukkovalur 	feud.Barony
	提卢瓦那摩罗Tiruvannamalai 	feud.Barony
}

func (c *多伽头罗TagadurCounty) BJinji旖耆() feud.Barony {
	return c.旖耆Jinji
}
    
func (c *多伽头罗TagadurCounty) BPadaividu波多毗都() feud.Barony {
	return c.波多毗都Padaividu
}
    
func (c *多伽头罗TagadurCounty) BSatyamangalam娑底也瞢伽罗() feud.Barony {
	return c.娑底也瞢伽罗Satyamangalam
}
    
func (c *多伽头罗TagadurCounty) BSiyyamangalam斯耶门伽愣() feud.Barony {
	return c.斯耶门伽愣Siyyamangalam
}
    
func (c *多伽头罗TagadurCounty) BTagadur多伽头罗() feud.Barony {
	return c.多伽头罗Tagadur
}
    
func (c *多伽头罗TagadurCounty) BTirukkovalur提卢伽婆卢() feud.Barony {
	return c.提卢伽婆卢Tirukkovalur
}
    
func (c *多伽头罗TagadurCounty) BTiruvannamalai提卢瓦那摩罗() feud.Barony {
	return c.提卢瓦那摩罗Tiruvannamalai
}
    
var CTagadur多伽头罗 TagadurCounty = &多伽头罗TagadurCounty{}

func init() {
	f := CTagadur多伽头罗.(*多伽头罗TagadurCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1120",
		Title:     "tagadur",
		TitleName: "多伽头罗",
		TitleCode: "c_tagadur",
		Baronies:  map[string]feud.Barony{},
	}

	f.旖耆Jinji = BJinji旖耆
	f.旖耆Jinji.SetParent(f)
	
	f.波多毗都Padaividu = BPadaividu波多毗都
	f.波多毗都Padaividu.SetParent(f)
	
	f.娑底也瞢伽罗Satyamangalam = BSatyamangalam娑底也瞢伽罗
	f.娑底也瞢伽罗Satyamangalam.SetParent(f)
	
	f.斯耶门伽愣Siyyamangalam = BSiyyamangalam斯耶门伽愣
	f.斯耶门伽愣Siyyamangalam.SetParent(f)
	
	f.多伽头罗Tagadur = BTagadur多伽头罗
	f.多伽头罗Tagadur.SetParent(f)
	
	f.提卢伽婆卢Tirukkovalur = BTirukkovalur提卢伽婆卢
	f.提卢伽婆卢Tirukkovalur.SetParent(f)
	
	f.提卢瓦那摩罗Tiruvannamalai = BTiruvannamalai提卢瓦那摩罗
	f.提卢瓦那摩罗Tiruvannamalai.SetParent(f)
	
}
