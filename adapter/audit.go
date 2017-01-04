/*
Copyright 2016 Crunchy Data Solutions, Inc.
Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at
http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package adapter

import (
	"log"
)

// Audit will create a adapter decorator with auditing concerns.
func Audit(l *log.Logger) Decorator {
	return func(c Adapter) Adapter {
		return AdapterFunc(func(r []byte, i int) error {
			l.Printf("msg len=%d\n", i)
			return c.Do(r, i)
		})
	}
}
