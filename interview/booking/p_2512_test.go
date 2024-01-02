package booking

import (
	"fmt"
	"testing"
)

func Test_topStudents(t *testing.T) {
	p := []string{"fkeofjpc", "qq", "iio"}
	n := []string{"jdh", "khj", "eget", "rjstbhe", "yzyoatfyx", "wlinrrgcm"}
	report := []string{"rjstbhe eget kctxcoub urrmkhlmi yniqafy fkeofjpc iio yzyoatfyx khj iio", "gpnhgabl qq qq fkeofjpc dflidshdb qq iio khj qq yzyoatfyx", "tizpzhlbyb eget z rjstbhe iio jdh jdh iptxh qq rjstbhe", "jtlghe wlinrrgcm jnkdbd k iio et rjstbhe iio qq jdh", "yp fkeofjpc lkhypcebox rjstbhe ewwykishv egzhne jdh y qq qq", "fu ql iio fkeofjpc jdh luspuy yzyoatfyx li qq v", "wlinrrgcm iio qq omnc sgkt tzgev iio iio qq qq", "d vhg qlj khj wlinrrgcm qq f jp zsmhkjokmb rjstbhe"}
	sid := []int{96537918, 589204657, 765963609, 613766496, 43871615, 189209587, 239084671, 908938263}
	k := 3
	fmt.Println(topStudents(p, n, report, sid, k))
}
