package main

import (
	"fmt"
	"math/rand"
)

type Math = int
type English = int
type Chinese = int

func main() {
	var mathScore int = 100
	var mathScore2 Math = 100

	mathScore2 = mathScore
	fmt.Println(mathScore2)
	getScoreOfStudent("")
	vg := &votGame{
		student: []*student{
			&student{name: fmt.Sprintf("%d", 1)},
			&student{name: fmt.Sprintf("%d", 2)},
		},
	}
	leader := vg.goRun()
	fmt.Println(leader)
	leader.Distribute()

	var stdXQ = &student{name: "小强"}
	var ldXQ Leader = Leader(*stdXQ)
	ldXQ.Distribute()
}

// 类型重命名，提高类型识别度
func getScoreOfStudent(name string) (Math, English, Chinese) {
	return 0, 0, 0
}

type votGame struct {
	student []*student
}

type student struct {
	name     string
	agree    int
	disagree int
}

func (g *votGame) goRun() *Leader {
	for _, item := range g.student {
		randInt := rand.Int()
		if randInt%2 == 0 {
			item.voteA(g.student[randInt%len(g.student)])
		} else {
			item.voteD(g.student[randInt%len(g.student)])
		}
		item.voteA(g.student[0])
	}
	maxScore := -1
	maxScoreIndex := -1
	for i, item := range g.student {
		if maxScore < item.agree {
			maxScore = item.agree
			maxScoreIndex = i
		}
	}
	if maxScore >= 0 { // 判断是否大于0，如果没有学生，index就是默认值 -1
		// 将student 转成 Leader 对象
		return (*Leader)(g.student[maxScoreIndex])
	}
	return nil
}

// type Leader struct {
// 	student
// }

// 使用类型重定义
type Leader student

// 使用嵌套对象定义（继承）方式定义班长
func (l *Leader) Distribute() {
	fmt.Println("发作业了")
}

type FooTestFuncRenderFine func()

func (f *FooTestFuncRenderFine) test111() {

}

func (std *student) voteA(target *student) {
	target.agree++
}

func (std *student) voteD(target *student) {
	target.disagree++
}
