// Copyright 2014-2016 Amazon.com, Inc. or its affiliates. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"). You may
// not use this file except in compliance with the License. A copy of the
// License is located at
//
//	http://aws.amazon.com/apache2.0/
//
// or in the "license" file accompanying this file. This file is distributed
// on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
// express or implied. See the License for the specific language governing
// permissions and limitations under the License.

package version

import (
	"fmt"

	ecsversion "github.com/aws/amazon-ecs-agent/ecs-version"
)

// String construct the version info of ecs-init
func String() string {
	dirtyMark := ""
	// TODO: Do we need the linke flags if we update version with every release?
	//
	// Note that GitDirty is defined via command-line linker flags, so
	// it must be a string rather than a bool
	if ecsversion.GitDirty == true {
		dirtyMark = "*"
	}
	return fmt.Sprintf("ecs-init version %s (%s%s)", ecsversion.Version, dirtyMark, ecsversion.GitShortHash)
}

// PrintVersion print out the
func PrintVersion() error {
	if ecsversion.Version == "" || ecsversion.GitShortHash == "" {
		return fmt.Errorf("Version info not set")
	}

	fmt.Println(String())

	return nil
}
