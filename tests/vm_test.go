// Copyright 2014 The go-wanliuno Authors
// This file is part of the go-wanliuno library.
//
// The go-wanliuno library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-wanliuno library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-wanliuno library. If not, see <http://www.gnu.org/licenses/>.

package tests

import (
	"testing"

	"github.com/wanliuno/go-wanliuno/core/vm"
)

func TestVM(t *testing.T) {
	t.Parallel()
	vmt := new(testMatcher)
	vmt.slow("^vmPerformance")
	vmt.fails("^vmSystemOperationsTest.json/createNameRegistrator$", "fails without parallel execution")

	vmt.walk(t, vmTestDir, func(t *testing.T, name string, test *VMTest) {
		withTrace(t, test.json.Exec.GasLimit, func(vmconfig vm.Config) error {
			return vmt.checkFailure(t, name+"/trie", test.Run(vmconfig, false))
		})
		withTrace(t, test.json.Exec.GasLimit, func(vmconfig vm.Config) error {
			return vmt.checkFailure(t, name+"/snap", test.Run(vmconfig, true))
		})
	})
}
