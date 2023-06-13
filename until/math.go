package until

//测试其他包使用gomodule方式 来引入类库

//取最大值
func Max(a, b int64) int64{
	if(a>b){
		return a
	}
	return b
}

//取最小值
func Min(a, b int64) int64{
	if(a>b){
		return b
	}
	return a
}