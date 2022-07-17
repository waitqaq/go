package main

import "testing"

func TestCase1Part1(t *testing.T) {
	inputRecord("王强", 0.38)
	inputRecord("王强", 0.38)
	rankOfWQ, fatRateOfWQ := getRank("王强")
	if rankOfWQ != 1 {
		// 标记为结束，程序继续运行
		//t.Fail()
		t.Fatalf("预期 王强 第一，但是得到：%d", rankOfWQ)
	}
	if fatRateOfWQ != 0.38 {
		t.Fatalf("预期王强的体脂是 0.32,但得到的是 :%f", fatRateOfWQ)
	}
}
func TestCase1(t *testing.T) {
	inputRecord("王强", 0.38)
	inputRecord("王强", 0.32)
	rankOfWQ, fatRateOfWQ := getRank("王强")
	if rankOfWQ != 1 {
		// 标记为结束，程序继续运行
		//t.Fail()
		t.Fatalf("预期 王强 第一，但是得到：%d", rankOfWQ)
	}
	if fatRateOfWQ != 0.32 {
		t.Fatalf("预期王强的体脂是 0.32,但得到的是 :%f", fatRateOfWQ)
	}
	inputRecord("李静", 0.28)
	inputRecord("李静", 0.32)
	rankOfLJ, fatRateOfLJ := getRank("李静")
	if rankOfLJ != 1 {
		t.Fatalf("预期 王强 第一，但是得到：%d", rankOfLJ)
	}
	if fatRateOfLJ != 0.32 {
		t.Fatalf("预期王强的体脂是 0.32,但得到的是 :%f", fatRateOfLJ)
	}

}

func TestCase2(t *testing.T) {
	inputRecord("李强", 0.38)
	inputRecord("王强", 0.38)
	inputRecord("李静", 0.28)

	rankOfLJ, fatRateOfLJ := getRank("李静")
	if rankOfLJ != 1 {
		t.Fatalf("预期 李静 第一，但是得到：%d", rankOfLJ)
	}
	if fatRateOfLJ != 0.28 {
		t.Fatalf("预期李静的体脂是 0.32,但得到的是 :%f", fatRateOfLJ)
	}

	rankOfWQ, fatRateOfWQ := getRank("王强")
	if rankOfWQ != 2 {
		t.Fatalf("预期 王强 第二，但是得到：%d", rankOfWQ)
	}
	if fatRateOfWQ != 0.38 {
		t.Fatalf("预期王强的体脂是 0.32,但得到的是 :%f", fatRateOfWQ)
	}

	rankOfLQ, fatRateOfLQ := getRank("李强")
	if rankOfLQ != 2 {
		t.Fatalf("预期 李强 第二，但是得到：%d", rankOfLQ)
	}
	if fatRateOfLQ != 0.38 {
		t.Fatalf("预期李强的体脂是 0.38,但得到的是 :%f", fatRateOfLQ)
	}
}

func TestCase3(t *testing.T) {
	inputRecord("王强", 0.32)
	inputRecord("张伟")
	rankOfWQ, fatRateOfWQ := getRank("王强")
	if rankOfWQ != 2 {
		// 标记为结束，程序继续运行
		//t.Fail()
		t.Fatalf("预期 王强 第一，但是得到：%d", rankOfWQ)
	}
	if fatRateOfWQ != 0.32 {
		t.Fatalf("预期王强的体脂是 0.32,但得到的是 :%f", fatRateOfWQ)
	}
}
