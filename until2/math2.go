package until2

import(
	"github.com/yunkaiyueming/modtest/util"
)

//测试util包的函数类库使用，自动调用比较函数
func AutoCompre(autoStr string, a,b int64) int64{
	if(autoStr=="min"){
		return util.Min(a,b)
	}
	return util.Max(a,b)
}