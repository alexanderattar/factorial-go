package factorial_test

import (
	"fmt"
	"math/big"
	"testing"

	"github.com/reservemedia/factorial-go"
)

func TestGenerateFactorial(t *testing.T) {
	var cases = []struct {
		num      int
		expected string
	}{
		{1, "1"},
		{2, "2"},
		{3, "6"},
		{4, "24"},
		{5, "120"},
		{50, "30414093201713378043612608166064768844377641568960512000000000000"},
		{100, "93326215443944152681699238856266700490715968264381621468592963895217599993229915608941463976156518286253697920827223758251185210916864000000000000000000000000"},
		{200, "788657867364790503552363213932185062295135977687173263294742533244359449963403342920304284011984623904177212138919638830257642790242637105061926624952829931113462857270763317237396988943922445621451664240254033291864131227428294853277524242407573903240321257405579568660226031904170324062351700858796178922222789623703897374720000000000000000000000000000000000000000000000000"},
	}

	for _, c := range cases {
		num := big.NewInt(int64(c.num))
		expected := c.expected

		t.Run(fmt.Sprintf("n=%d", num), func(t *testing.T) {
			t.Parallel()

			got := factorial.GenerateFactorial(num).String()
			if expected != got {
				t.Errorf("expected %d to equal %s, got %s\n",
					num,
					expected,
					got,
				)
			}
		})
	}
}
