// Copyright 2021 Matrix Origin
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package overload

import (
	"github.com/matrixorigin/matrixone/pkg/container/types"
	"github.com/matrixorigin/matrixone/pkg/container/vector"
	"github.com/matrixorigin/matrixone/pkg/vm/process"
)

func UpdateEval(typ, toTyp types.T, c bool, v *vector.Vector, p *process.Process) (*vector.Vector, error) {
	var err error
	v, err = BinaryEval(Typecast, typ, toTyp, c, false, v, vector.New(types.Type{Oid: toTyp}), p)
	if err != nil {
		return nil, err
	}
	return v, nil
}