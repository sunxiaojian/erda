// Copyright (c) 2021 Terminus, Inc.
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

package ctxhelper

import (
	"context"
	"sync"

	"github.com/erda-project/erda/internal/apps/ai-proxy/vars"
	"github.com/erda-project/erda/pkg/reverseproxy"
)

func GetIsStream(ctx context.Context) (bool, bool) {
	value, ok := ctx.Value(reverseproxy.CtxKeyMap{}).(*sync.Map).Load(vars.MapKeyIsStream{})
	if !ok || value == nil {
		return false, false
	}
	isStream, ok := value.(bool)
	if !ok {
		return false, false
	}
	return isStream, true
}

func PutIsStream(ctx context.Context, isStream bool) {
	m := ctx.Value(reverseproxy.CtxKeyMap{}).(*sync.Map)
	m.Store(vars.MapKeyIsStream{}, isStream)
}
