package calc

func CalcBMI(tall float64, weight float64) (bmi float64) {
	if tall <= 0 {
		panic("身高不能是 0 或者负数")
	}
	// todo 验证体重合法性
	return weight / (tall * tall)
}
