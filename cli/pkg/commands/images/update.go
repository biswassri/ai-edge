/*
Copyright 2024. Open Data Hub Authors

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

package images

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/opendatahub-io/ai-edge/cli/pkg/commands/common"
	"github.com/opendatahub-io/ai-edge/cli/pkg/commands/flags"
)

var updateCmd = common.NewCmd(
	fmt.Sprintf(
		"update -%s model-id -%s version [-%s model-registry-url] [-%s params-file]",
		flags.FlagModelID.Shorthand(),
		flags.FlagVersionName.Shorthand(),
		flags.FlagModelRegistryURL.Shorthand(),
		flags.FlagParams.Shorthand(),
	),
	"Update image parameters for an edge model in the model registry.",
	`Update image parameters for an edge model in the model registry.

This command allows you to update the build parameters stored in the model registry for a specific version of a model.
`,
	cobra.NoArgs,
	[]flags.Flag{
		flags.FlagModelID.SetRequired(),
		flags.FlagVersionName.SetRequired(),
		flags.FlagModelRegistryURL.SetParentFlag(),
		flags.FlagParams,
	},
	common.SubCommandUpdate,
	NewImagesModel,
)
