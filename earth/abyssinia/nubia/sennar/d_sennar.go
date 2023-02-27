package sennar

import (
	"github.com/thalesfu/CK2Commands/earth/abyssinia/nubia/sennar/alodia"
	"github.com/thalesfu/CK2Commands/earth/abyssinia/nubia/sennar/bayuda"
	"github.com/thalesfu/CK2Commands/earth/abyssinia/nubia/sennar/darfur"
	"github.com/thalesfu/CK2Commands/earth/abyssinia/nubia/sennar/kosti"
	"github.com/thalesfu/CK2Commands/earth/abyssinia/nubia/sennar/sennar"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type SennarDuke interface {
    feud.Duke
    CAlodia阿洛地亚() 	alodia.AlodiaCounty
    CBayuda拜尤达() 	bayuda.BayudaCounty
    CDarfur达尔富尔() 	darfur.DarfurCounty
    CKosti科斯提() 	kosti.KostiCounty
    CSennar森纳尔() 	sennar.SennarCounty
}

type 阿洛地亚SennarDuke struct {
	feud.BaseDuke
	阿洛地亚Alodia 	alodia.AlodiaCounty
	拜尤达Bayuda 	bayuda.BayudaCounty
	达尔富尔Darfur 	darfur.DarfurCounty
	科斯提Kosti 	kosti.KostiCounty
	森纳尔Sennar 	sennar.SennarCounty
}

func (d *阿洛地亚SennarDuke) CAlodia阿洛地亚() alodia.AlodiaCounty {
	return d.阿洛地亚Alodia
}
    
func (d *阿洛地亚SennarDuke) CBayuda拜尤达() bayuda.BayudaCounty {
	return d.拜尤达Bayuda
}
    
func (d *阿洛地亚SennarDuke) CDarfur达尔富尔() darfur.DarfurCounty {
	return d.达尔富尔Darfur
}
    
func (d *阿洛地亚SennarDuke) CKosti科斯提() kosti.KostiCounty {
	return d.科斯提Kosti
}
    
func (d *阿洛地亚SennarDuke) CSennar森纳尔() sennar.SennarCounty {
	return d.森纳尔Sennar
}
    
var DSennar阿洛地亚 SennarDuke = &阿洛地亚SennarDuke{}

func init() {
	f := DSennar阿洛地亚.(*阿洛地亚SennarDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "sennar",
		TitleName: "阿洛地亚",
		TitleCode: "d_sennar",
		Counties:  map[string]feud.County{},
	}

	f.阿洛地亚Alodia = alodia.CAlodia阿洛地亚
	f.阿洛地亚Alodia.SetParent(f)
	
	f.拜尤达Bayuda = bayuda.CBayuda拜尤达
	f.拜尤达Bayuda.SetParent(f)
	
	f.达尔富尔Darfur = darfur.CDarfur达尔富尔
	f.达尔富尔Darfur.SetParent(f)
	
	f.科斯提Kosti = kosti.CKosti科斯提
	f.科斯提Kosti.SetParent(f)
	
	f.森纳尔Sennar = sennar.CSennar森纳尔
	f.森纳尔Sennar.SetParent(f)
	
}
