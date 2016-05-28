/*
Copyright 2016 The Kubernetes Authors All rights reserved.

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

package cluster

import (
	gflag "flag"
	"fmt"
	"strings"

	"k8s.io/minikube/pkg/minikube/constants"
)

const (
	remoteLocalKubeErrPath = "/var/log/localkube.err"
	remoteLocalKubeOutPath = "/var/log/localkube.out"
)

// Kill any running instances.
var stopCommand = "sudo killall localkube | true"

var startCommandFmtStr = `
# Run with nohup so it stays up. Redirect logs to useful places.
PATH=/usr/local/sbin:$PATH nohup sudo /usr/local/bin/localkube start %s --generate-certs=false --logtostderr=true > %s 2> %s < /dev/null &
`

var logsCommand = fmt.Sprintf("tail -n +1 %s %s", remoteLocalKubeErrPath, remoteLocalKubeOutPath)

func GetStartCommand() string {
	flagVals := make([]string, len(constants.LogFlags))
	for _, logFlag := range constants.LogFlags {
		if logVal := gflag.Lookup(logFlag); logVal != nil && logVal.Value.String() != logVal.DefValue {
			flagVals = append(flagVals, fmt.Sprintf("--%s %s", logFlag, logVal.Value.String()))
		}
	}
	flags := strings.Join(flagVals, " ")
	return fmt.Sprintf(startCommandFmtStr, flags, remoteLocalKubeErrPath, remoteLocalKubeOutPath)
}