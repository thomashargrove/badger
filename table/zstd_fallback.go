// +build !cgo

/*
 * Copyright 2019 Dgraph Labs, Inc. and Contributors
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package table

import (
	"errors"
)

var errZstdCgo = errors.New("zstd compression requires building with cgo enabled")

func zstdDecompress(dst, src []byte) ([]byte, error) {
	return nil, errZstdCgo
}

func zstdCompress(dst, src []byte) ([]byte, error) {
	return nil, errZstdCgo
}
