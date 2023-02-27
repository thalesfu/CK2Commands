package bosnia

import (
	"github.com/thalesfu/CK2Commands/earth/byzantium/croatia/bosnia/bilino_polje"
	"github.com/thalesfu/CK2Commands/earth/byzantium/croatia/bosnia/rama"
	"github.com/thalesfu/CK2Commands/earth/byzantium/croatia/bosnia/soli"
	"github.com/thalesfu/CK2Commands/earth/byzantium/croatia/bosnia/usora"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type BosniaDuke interface {
    feud.Duke
    CBilino_polje弗尔赫波斯纳() 	bilino_polje.Bilino_poljeCounty
    CRama索利() 	rama.RamaCounty
    CSoli乌索拉() 	soli.SoliCounty
    CUsora下波斯尼亚() 	usora.UsoraCounty
}

type 波斯尼亚BosniaDuke struct {
	feud.BaseDuke
	弗尔赫波斯纳Bilino_polje 	bilino_polje.Bilino_poljeCounty
	索利Rama 	rama.RamaCounty
	乌索拉Soli 	soli.SoliCounty
	下波斯尼亚Usora 	usora.UsoraCounty
}

func (d *波斯尼亚BosniaDuke) CBilino_polje弗尔赫波斯纳() bilino_polje.Bilino_poljeCounty {
	return d.弗尔赫波斯纳Bilino_polje
}
    
func (d *波斯尼亚BosniaDuke) CRama索利() rama.RamaCounty {
	return d.索利Rama
}
    
func (d *波斯尼亚BosniaDuke) CSoli乌索拉() soli.SoliCounty {
	return d.乌索拉Soli
}
    
func (d *波斯尼亚BosniaDuke) CUsora下波斯尼亚() usora.UsoraCounty {
	return d.下波斯尼亚Usora
}
    
var DBosnia波斯尼亚 BosniaDuke = &波斯尼亚BosniaDuke{}

func init() {
	f := DBosnia波斯尼亚.(*波斯尼亚BosniaDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "bosnia",
		TitleName: "波斯尼亚",
		TitleCode: "d_bosnia",
		Counties:  map[string]feud.County{},
	}

	f.弗尔赫波斯纳Bilino_polje = bilino_polje.CBilino_polje弗尔赫波斯纳
	f.弗尔赫波斯纳Bilino_polje.SetParent(f)
	
	f.索利Rama = rama.CRama索利
	f.索利Rama.SetParent(f)
	
	f.乌索拉Soli = soli.CSoli乌索拉
	f.乌索拉Soli.SetParent(f)
	
	f.下波斯尼亚Usora = usora.CUsora下波斯尼亚
	f.下波斯尼亚Usora.SetParent(f)
	
}
